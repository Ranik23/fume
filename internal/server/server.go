package server

import (
	"context"
	"fume/config"
	sl "fume/internal/lib/logger"
	"fume/internal/repository"
	"fume/internal/server/handlers"
	"fume/internal/server/middlewares"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/gorilla/mux"
)

type Server struct {
	Repository  *repository.Repository
	Address    	string
	Logger     	*slog.Logger
	Router     	*mux.Router
	httpServer 	*http.Server
	Sig        	chan os.Signal
}

func NewServer(config *config.Config, repository *repository.Repository) *Server {
	return &Server {
		Address: 	config.ServerConfig.Host + ":" + config.ServerConfig.Port,
		Logger:  	repository.Logger,
		Router:  	mux.NewRouter(),
		Sig:     	make(chan os.Signal, 1),
		Repository: repository,
	}
}
func (s *Server) handle() {
	fs := http.FileServer(http.Dir("/home/anton/rosl/internal/static"))

	s.Router.Handle("/static/", http.StripPrefix("/static/", fs))

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	s.Router.HandleFunc("/login", handlers.ShowLogin(s.Repository)).Methods("GET")
	s.Router.HandleFunc("/login", handlers.HandleLogin(s.Repository)).Methods("POST")
	s.Router.HandleFunc("/submit", handlers.SubmitHandler(s.Repository))
	s.Router.HandleFunc("/config", handlers.ConfigHandler(s.Repository))

	s.Router.Use(middlewares.TimeMiddleware)
}

func (s *Server) Run() error {

	s.handle()

	s.httpServer = &http.Server {
		Addr:    s.Address,
		Handler: s.Router,
	}
	signal.Notify(s.Sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		s.Logger.Info("Starting server", "address", s.Address)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error("Error starting server: ", sl.Err(err))
		}
	}()

	<-s.Sig

	s.Logger.Info("Signal to stop has been caught")

	return s.Stop(context.Background())
}

func (s *Server) Stop(ctx context.Context) error {
	s.Logger.Info("Shutting down server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.Logger.Error("Failed to gracefully shutdown server", sl.Err(err))
		return err
	}
	s.Logger.Info("Server gracefully stopped.")
	return nil
}
