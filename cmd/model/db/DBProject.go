package db

import (
	"RbsBurndownChart/cmd/types"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

//Объект обращения к базе данных, параметр connSpec задан объектом Config пакета config
type DBProject struct {
	connSpec string
	db       *pgx.Conn
}

//NewDBProject() инициализирует DBProject, и возвращает указатель на него.
//Аргумент connSpec задаёт адрес для соединения с базой данных
func NewDBProject(connSpec string) *DBProject {
	return &DBProject{}
}

// Read() реализует чтение данных по конкретному проекту исходя из параметра логина projectId, заданных
// в структуре types.Project
func (d *DBProject) Read(projectId int64) (*types.Project, error) {
	return pgxConnProject(projectId)
}

//pgxConnProject() - функция принимает аргумент табличного идентификатора реализует внутри себя объект типа *types.Project обращается к базе данных
//в соответствии с заданным запросом, заполняет и возвращает объект и ошибку
func pgxConnProject(arg int64) (*types.Project, error) {
	//Объявляем пустую структуру types.Project а также переменные которые примут значение из базы данных
	var id int64
	var stringTasksSource string
	resultStruct := types.Project{}
	//Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://admin:1234@localhost:5432/burndown_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	//Считываем необходимые параметры для всей заданной таблицы, проверяем ошибки
	err = conn.QueryRow(context.Background(), "SELECT proj.c_id, proj.c_name, CAST(proj.c_sprint_start_date AS TEXT), tref.c_name, proj.c_task_list_file_path FROM public.t_project AS proj LEFT JOIN t_ref_tasks_source AS tref ON proj.fk_tasks_source = tref.c_id WHERE proj.c_id=$1", arg).Scan(&id, &resultStruct.Name, &resultStruct.SprintStartDate, &stringTasksSource, &resultStruct.TaskListFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow in Project failed: %v\n", err)
		os.Exit(1)
	}
	//Для приведения типа данных полученного значения атрибута stringTasksSource к типу int, выполняем условие:
	switch stringTasksSource {
	case "File":
		resultStruct.TasksSource = types.File
	case "Data_base":
		resultStruct.TasksSource = types.DB
	case "Web":
		resultStruct.TasksSource = types.Service
	}
	//Заполняем структуру полученными значениями
	resultStruct = types.Project{
		Id:               id,
		Name:             resultStruct.Name,
		SprintStartDate:  resultStruct.SprintStartDate,
		TasksSource:      resultStruct.TasksSource,
		TaskListFilePath: resultStruct.TaskListFilePath,
	}
	return &resultStruct, err
}
