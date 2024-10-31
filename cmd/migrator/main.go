package main

import (
	"context"
	sl "fume/internal/lib/logger"
	"fume/internal/storage/ent"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func main() {

	Logger := sl.SetUpLogger("local")

    driver, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
    if err != nil {
        Logger.Error("failed opening connection to database: %v", sl.Err(err))
    }
    client := ent.NewClient(ent.Driver(driver))
    
    if err := client.Schema.Create(context.Background()); err != nil {
        Logger.Error("failed to create schema resources: %v", sl.Err(err))
    }

	Logger.Info("migrations applied succesfully")
}
