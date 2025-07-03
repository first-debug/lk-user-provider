package app

import (
	"main/internal/config"
	"main/internal/database"
	"main/internal/server"
)

type App struct {
	server *server.Server
	config *config.Config
}

func New(config *config.Config) *App {
	db := database.GetDB()
	srv := server.NewServer(db)
	return &App{
		server: srv,
		config: config,
	}
}
func (a *App) Run() error {
	a.server.Start(a.config.URL+a.config.Port, a.config.Port)
	return nil
}
