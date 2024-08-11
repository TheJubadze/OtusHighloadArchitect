package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/TheJubadze/OtusHighloadArchitect/peepl/config"
	"github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/app"
	"github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/logger"
	"github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/model"
	"github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/storage"
	"github.com/TheJubadze/OtusHighloadArchitect/peepl/utils"
)

type HttpServer struct {
	host    string
	port    int
	storage storage.Storage
	server  *http.Server
}

func NewHttpServer(app *app.App) *HttpServer {
	s := &HttpServer{
		host:    config.Config.HttpServer.Host,
		port:    config.Config.HttpServer.Port,
		storage: app.Storage(),
	}

	loggedMux := LoggingMiddleware(s.routes())

	s.server = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", s.host, s.port),
		Handler:           loggedMux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadTimeout:       15 * time.Second,
	}
	return s
}

func (s *HttpServer) Start() error {
	logger.Log.Printf("Starting HTTP server at %s:%d", s.host, s.port)
	return s.server.ListenAndServe()
}

func (s *HttpServer) Stop(ctx context.Context) error {
	logger.Log.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}

func (s *HttpServer) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", registerUser(s.storage))
	mux.HandleFunc("POST /login", loginUser(s.storage))
	return mux
}

func registerUser(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error hashing password: %s", err), http.StatusInternalServerError)
			return
		}

		user.Password = string(hashedPassword)

		// Add user to the database
		if err := s.AddUser(user); err != nil {
			http.Error(w, fmt.Sprintf("Error adding user: %s", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func loginUser(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Find user by firstname
		user, err := s.GetUser(credentials.Login)
		if err != nil {
			http.Error(w, fmt.Sprintf("User not found: %s", err), http.StatusNotFound)
			return
		}

		// Check the password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
			http.Error(w, fmt.Sprintf("Invalid password: %s", err), http.StatusUnauthorized)
			return
		}

		// Create JWT token
		tokenString, err := utils.GenerateJWT(user.Login)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Send token as response
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: tokenString,
		})

		w.WriteHeader(http.StatusOK)
	}
}
