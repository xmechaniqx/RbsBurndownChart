package db

import (
	"RbsBurndownChart/cmd/types"

	"github.com/jackc/pgx"
)

type DBProject struct {
	connSpec string
	db       *pgx.Conn
}

func NewDBProject(connSpec string) *DBProject {
	return &DBProject{}
}

func (db *DBProject) Read(projectId int64) (*types.Developer, error) {
	return nil, nil
}
