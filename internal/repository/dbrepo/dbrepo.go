package dbrepo

import (
	"database/sql"
	"github.com/dapetoo/go-bookings/internal/config"
	"github.com/dapetoo/go-bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo creates a new repository
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

// NewPostgresRepo creates a new repository
func NewTestingsRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
		DB:  conn,
	}
}
