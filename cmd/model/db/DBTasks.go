package db

import (
	"RbsBurndownChart/cmd/types"

	"github.com/jackc/pgx"
)

type DBTasks struct {
	connSpec string
	db       *pgx.Conn
}

func NewDBTasks(connSpec string) *DBTasks {
	return &DBTasks{}
}
func (tr *DBTasks) Read(p types.Project) ([]types.Task, error) {
	mockTasks := []types.Task{
		types.Task{
			Title: "Dummy BD task 1",
			Cost:  5,
		},
	}
	return mockTasks, nil
}
