package server

import (
	"fume/internal/server/handlers"
	"fume/internal/server/middlewares"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)



type Server struct {
	Address string
	Logger *slog.Logger
	Router *mux.Router
}


func NewServer(address string) *Server {
	return &Server{
		Address: address,
		Logger: slog.Default(),
		Router : mux.NewRouter(),
	}
}

func (s *Server) Handle() {
	s.Router.HandleFunc("/", handlers.HomeHandler)
	s.Router.Use(middlewares.TimeMiddleware)
}


func (s *Server) Run() {
	go func() {
		http.ListenAndServe(s.Address, s.Router);
		s.Logger.Info("Listening on ", s.Address)
	}()
}