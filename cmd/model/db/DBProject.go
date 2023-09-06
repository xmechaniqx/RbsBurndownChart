package db

import (
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
