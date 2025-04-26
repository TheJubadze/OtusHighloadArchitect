package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/config"
	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/app"
	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/logger"
	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/server/http"
	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/storage"

	_ "github.com/lib/pq"
)

var (
	configFile = flag.String("config", "/etc/peepl/config.yaml", "Path to configuration file")
	cfg        = Config
)

func init() {
	flag.Parse()
}

func main() {
	if err := Init(*configFile); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	SetupLogger(cfg.Logger.Level)

	storage, closeFunc, err := initStorage()
	defer closeFunc()
	if err != nil {
		Log.Fatalf("Error initializing storage: %s", err)
	}

	app := NewApp(storage)
	server := NewHttpServer(app)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			Log.Error("failed to stop server: " + err.Error())
		}
	}()

	if err := server.Start(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			Log.Fatalf("Error starting server: %s", err)
		}
	}
}

func initStorage() (Storage, func(), error) {
	var closeFunc func()
	storage, err := NewSqlStorage(cfg.Storage.DSN)
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing sql storage: %w", err)
	}

	Log.Println("connected to psql")

	err = storage.Migrate(cfg.Storage.MigrationsDir)
	if err != nil {
		return nil, nil, fmt.Errorf("error migrating sql storage: %w", err)
	}

	Log.Println("migrated psql")

	closeFunc = func() {
		if err := storage.Close(); err != nil {
			Log.Println("cannot close psql connection", err)
		}
	}

	return storage, closeFunc, nil
}
