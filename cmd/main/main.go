package main

import (
	"fume/config"
	sl "fume/internal/lib/logger"
	"fume/internal/repository"
	"fume/internal/server"
	"fume/internal/storage"
	_ "github.com/lib/pq"
)

func main() {

	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	Logger := sl.SetUpLogger("local")

	Config, err := config.GetConfig("/home/anton/rosl/config/config.yaml")
	if err != nil {
		Logger.Error("failed to get the config", sl.Err(err))
		return
	}

	DataBase := storage.NewStorage(dsn, Logger)
	
	Repository, err := repository.NewRepository(Config, Logger, DataBase)
	if err != nil {
		Logger.Error("failed to create the repository", sl.Err(err))
		return
	}

	srv := server.NewServer(Config, Repository)

	if err := srv.Run(); err != nil {
		Logger.Error("Failed to start server: ", sl.Err(err))
	}
}