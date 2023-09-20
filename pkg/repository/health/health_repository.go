package health

import (
	"context"
	"database/sql"
)

type Repository struct {
	client *sql.DB
}

func NewRepository(client *sql.DB) Repository {
	return Repository{client: client}
}

func (h Repository) IsHealthy() bool {
	// Use PingContext to check the database client's health
	ctx := context.Background()
	if err := h.client.PingContext(ctx); err != nil {
		return false
	}
	return true
}
