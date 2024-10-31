package storage

import (
	"context"
	sl "fume/internal/lib/logger"
	"fume/internal/storage/ent"
	"log/slog"
	"net/http"
	"time"
)

type Storage struct {
	Logger 	 *slog.Logger
	Client	 *ent.Client
}

func NewStorage(dsn string, logger *slog.Logger) *Storage {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		logger.Error("failed to create the client", sl.Err(err))
		return nil
	}

	return &Storage {
		Logger: logger,
		Client: client,
	}
}

func (s *Storage) CreateRequest(request *http.Request) error {
	ctx := context.Background()

	_, err := s.Client.Request.
		Create().
		SetMethod(request.Method).
		SetEndpoint(request.URL.Path).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetID(1).
		Save(ctx)

	if err != nil {
		s.Logger.Error("failed to create request", sl.Err(err))
		return err
	}

	s.Logger.Info("request created successfully")

	return nil
}

func (s *Storage) DeleteRequest(id int) error {
	ctx := context.Background()

	err := s.Client.Request.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		s.Logger.Error("Failed to delete request from database", sl.Err(err))
		return err
	}
	return nil
}

func (s *Storage) GetRequest(id int) (*ent.Request, error) {
	ctx := context.Background()

	request, err := s.Client.Request.Get(ctx, id)
	if err != nil {
		s.Logger.Error("Failed to get request from database", sl.Err(err))
		return nil, err
	}

	return request, nil
}


func (s *Storage) Clear(ctx context.Context) error {
	n, err := s.Client.Request.Delete().Exec(ctx);
	if err != nil {
		s.Logger.Error("Failed to clear the database", sl.Err(err))
		return err
	}
	s.Logger.Info("Vertices deleted", "count" , n)

	return nil
}







