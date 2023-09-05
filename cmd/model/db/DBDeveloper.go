package db

import (
	"RbsBurndownChart/cmd/types"

	"github.com/jackc/pgx"
)

type DBDeveloper struct {
	connSpec string
	db       *pgx.Conn
}

func NewDBDeveloper(connSpec string) *DBDeveloper {
	return &DBDeveloper{}
}

func (db *DBDeveloper) Read(devLogin string) (*types.Developer, error) {
	return nil, nil
}
func (db *DBDeveloper) ReadAll(projectId int64) ([]types.Developer, error) {
	return nil, nil
}
