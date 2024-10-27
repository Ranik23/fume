package server

import (
	"context"
	//"fume/internal/server/handlers"
	"fume/internal/server/middlewares"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/gorilla/mux"
	"github.com/lmittmann/tint"
)

type Server struct {
	Address    string
	Logger     *slog.Logger
	Router     *mux.Router
	httpServer *http.Server
	sig        chan os.Signal
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
		Logger:  slog.New(tint.NewHandler(os.Stdout, nil)),
		Router:  mux.NewRouter(),
		sig:     make(chan os.Signal, 1), 
	}
}

func (s *Server) registerRoutes() {
	fs := http.FileServer(http.Dir("/home/anton/rosl/internal/static"))
    s.Router.Handle("/static/", http.StripPrefix("/static/", fs))

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "/home/anton/rosl/internal/static/index.html")
    })
}

func (s *Server) registerMiddlewares() {
	s.Router.Use(middlewares.TimeMiddleware)
}

func (s *Server) Handle() {
	s.registerRoutes()
	s.registerMiddlewares()
}

func (s *Server) Run() error {

	s.httpServer = &http.Server{
		Addr:    s.Address,
		Handler: s.Router,
	}

	signal.Notify(s.sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		s.Logger.Info("Starting server on ", s.Address)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error("Error starting server: ", err)
		}
	}()

	<-s.sig

	return s.Stop(context.Background())
}

func (s *Server) Stop(ctx context.Context) error {
	s.Logger.Info("Shutting down server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.Logger.Error("Failed to gracefully shutdown server: ", err)
		return err
	}
	s.Logger.Info("Server gracefully stopped.")
	return nil
}
