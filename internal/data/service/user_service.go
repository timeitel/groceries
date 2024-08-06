package service

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/go-libsql"
)

type service struct {
	repo UserRepository
}

func (s *service) GetUserItems(c echo.Context) error {

}

func New(db *sql.DB) service {
	return service{
		conn: db,
	}
}
