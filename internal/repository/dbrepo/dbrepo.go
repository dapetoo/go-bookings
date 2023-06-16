package dbrepo

import (
	"database/sql"
	"github.com/dapetoo/go-bookings/internal/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo creates a new repository
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) *postgresDBRepo {
	return &postgresDBRepo{
		DB:  conn,
		App: a,
	}
}
