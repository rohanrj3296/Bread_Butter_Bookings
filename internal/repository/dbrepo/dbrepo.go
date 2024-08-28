package dbrepo

import (
	"database/sql"

	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/repository"
	"github.com/rohanrj3296/Bread_Butter_Bookings/packages/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo{
	return &postgresDBRepo{
		App: a,
		DB:conn,
	}
}