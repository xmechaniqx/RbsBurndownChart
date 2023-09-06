package db

import (
	"RbsBurndownChart/cmd/types"

	"github.com/jackc/pgx"
)

//Объект обращения к базе данных, параметр connSpec задан объектом Config пакета config.
type DBTasks struct {
	connSpec string
	db       *pgx.Conn
}

// NewDBTasks() инициализирует DBTasks, и возвращает указатель на него.
// Аргумент connSpec задаёт адрес для соединения с базой данных.
func NewDBTasks(connSpec string) *DBTasks {
	return &DBTasks{}
}

// Read() реализует чтение списка задач для проекта, заданных
// в структуре types.Task.
func (tr *DBTasks) Read(p types.Project) ([]types.Task, error) {
	mockTasks := []types.Task{
		types.Task{
			Title: "Dummy BD task 1",
			Cost:  5,
		},
	}
	return mockTasks, nil
}
