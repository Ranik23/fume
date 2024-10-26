package server

import (
	"log/slog"
	"net/http"
	"github.com/gorilla/mux"
)



type Server struct {
	Address string
	Logger *slog.Logger
}



func NewServer(address string) *Server {
	return &Server{
		Address: address,
		Logger: slog.Default(),
	}
}


func (s *Server) Run() {

	mux := mux.NewRouter()

	go func() {
		http.ListenAndServe(s.Address, mux)
	}()
}