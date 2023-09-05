package db

import "github.com/jackc/pgx"

type DBTasks struct {
	connSpec string
	db       *pgx.Conn
}

func NewDBTasks(connSpec string) *DBTasks {
	return &DBTasks{}
}
