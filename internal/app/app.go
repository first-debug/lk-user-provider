package app

import (
	"context"
	"errors"
	"log/slog"
	"main/internal/config"
	"main/internal/database"
	"main/internal/server"
	"sync"
	"sync/atomic"
)

type App struct {
	log         *slog.Logger
	server      *server.Server
	cfg         *config.Config
	userStorage database.UserStorage
}

func New(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, log *slog.Logger, isShutDown *atomic.Bool) (*App, error) {
	userStorage := database.NewMySQLUserStorage(cfg.DB_URL)
	srv := server.NewServer(ctx, log, isShutDown, userStorage)
	return &App{
		log:         log,
		server:      srv,
		cfg:         cfg,
		userStorage: userStorage,
	}, nil
}

func (a *App) Run() error {
	a.log.Info("Запуск HTTP сервера по адресу: '" + a.cfg.URL + ":" + a.cfg.Port + "'...")
	a.server.Start(a.cfg.Env, a.cfg.URL+":"+a.cfg.Port)
	return nil
}

func (a *App) ShutDown(shutDownCtx context.Context) error {
	if a == nil {
		return errors.New("app is nil")
	}

	err := errors.Join(
		a.server.ShutDown(shutDownCtx),
	)

	return err

}
