package db

import (
	"RbsBurndownChart/cmd/types"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

//Объект обращения к базе данных, параметр connSpec задан объектом Config пакета config.
type DBTasks struct {
	connSpec string
	db       *pgx.Conn
}

//NewDBTasks() инициализирует DBTasks, и возвращает указатель на него.
//Аргумент connSpec задаёт адрес для соединения с базой данных.
func NewDBTasks(connSpec string) *DBTasks {
	return &DBTasks{}
}

// Read() реализует чтение списка задач для проекта, заданных
// в структуре types.Task.
func (tr *DBTasks) Read(p types.Project) ([]types.Task, error) {
	//Объявляем пустую структуру и массив структур types.Task а также переменные которые примут значение из базы данных
	taskStruct := types.Task{}
	resultTask := []types.Task{}
	var title string
	var cost int64
	//Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://admin:1234@localhost:5432/burndown_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	//Считываем необходимые параметры для всей заданной таблицы
	all, err := conn.Query(context.Background(), "SELECT t_task.c_title, t_task.c_cost FROM t_task WHERE fk_project =$1", p.Id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	//Цикл чтения строк таблицы и запись данных в структуру taskStruct, далее заполненяем массив объектов полученной структурой
	for all.Next() {
		err := all.Scan(&title, &cost)
		if err != nil {
			fmt.Println("next error")
		}
		if all.Err() != nil {
			fmt.Println("scan error")
		}
		taskStruct = types.Task{
			Title: title,
			Cost:  cost,
		}
		resultTask = append(resultTask, taskStruct)
	}
	return resultTask, err
}
