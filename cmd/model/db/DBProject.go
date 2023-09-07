package db

import (
	"RbsBurndownChart/cmd/types"

	"github.com/jackc/pgx"
)

//Объект обращения к базе данных, параметр connSpec задан объектом Config пакета config.
type DBProject struct {
	connSpec string
	db       *pgx.Conn
}

// NewDBProject() инициализирует DBProject, и возвращает указатель на него.
// Аргумент connSpec задаёт адрес для соединения с базой данных.
func NewDBProject(connSpec string) *DBProject {
	return &DBProject{}
}
func (d *DBProject) Read(projectId int64) (*types.Project, error) {
	mockDB := types.Project{
		Name:             "front",
		SprintStartDate:  "22-08-2023",
		TasksSource:      1,
		TaskListFilePath: "",
	}
	return &mockDB, nil
}
