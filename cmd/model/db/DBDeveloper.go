package db

import (
	"RbsBurndownChart/cmd/types"

	"github.com/jackc/pgx"
)

//Объект обращения к базе данных, параметр connSpec задан объектом Config пакета config.
type DBDeveloper struct {
	connSpec string
	db       *pgx.Conn
}

// NewDBDeveloper() инициализирует DBDeveloper, и возвращает указатель на него.
// Аргумент connSpec задаёт адрес для соединения с базой данных.
func NewDBDeveloper(connSpec string) *DBDeveloper {
	return &DBDeveloper{}
}

// Read() реализует чтение данных по конкретному разработчику исходя из параметра логина devLogin, заданных
// в структуре types.Developer.
func (d *DBDeveloper) Read(devLogin string) (*types.Developer, error) {
	test := make([]int64, 0)
	test = append(test, 1, 5) //ID номер проекта в базе
	mockDB := types.Developer{
		Id:       1,
		Projects: test,
		FullName: "Petrov Ivan", //first name+last name
		Login:    "ipetrov",
		Position: "developer",
	}
	return &mockDB, nil
}

// ReadAll() реализует чтение списка разработчиков для проекта (projectId) из базы данных, заданной
// в структуре DBDeveloper.
func (db *DBDeveloper) ReadAll(projectId int64) ([]types.Developer, error) {
	return nil, nil
}
