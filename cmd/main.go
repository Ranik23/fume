package main

import (
	"fume/internal/server"
)



func main() {

	srv := server.NewServer("localhost:8080")

	srv.Handle()

	if err := srv.Run(); err != nil {
		srv.Logger.Error("Failed to start server: ", "error", err.Error())
	}
}