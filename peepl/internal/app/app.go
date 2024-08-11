package app

import (
	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/storage"
)

type App struct {
	storage Storage
}

func NewApp(storage Storage) *App {
	return &App{storage: storage}
}

func (a *App) Storage() Storage {
	return a.storage
}
