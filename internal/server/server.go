package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
)

type Server struct {
	ctx            context.Context
	log            *slog.Logger
	isShuttingDown *atomic.Bool
	router         *http.ServeMux
	server         http.Server
}

func NewServer(ctx context.Context, log *slog.Logger, isShuttingDown *atomic.Bool) *Server {
	s := &Server{
		ctx:            ctx,
		log:            log,
		isShuttingDown: isShuttingDown,
		router:         http.NewServeMux(),
	}
	s.router.HandleFunc("GET /ping", s.handlePing)
	return s
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}

func (s *Server) Start(env, addr string) {

	s.server = http.Server{
		Addr:    addr,
		Handler: s.router,
		BaseContext: func(_ net.Listener) context.Context {
			return s.ctx
		},
	}
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.log.Error("failed to start server", "error", err)
	}
}

func (s *Server) ShutDown(shutDownCtx context.Context) error {
	return s.server.Shutdown(shutDownCtx)
}
