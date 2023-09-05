package db

import (
	"github.com/jackc/pgx"
)

type DBProject struct {
	connSpec string
	db       *pgx.Conn
}

func NewDBProject(connSpec string) *DBProject {
	return &DBProject{}
}
