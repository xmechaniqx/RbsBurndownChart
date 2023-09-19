package db

import (
	"RbsBurndownChart/cmd/types"
	"context"
	"fmt"

	"github.com/jackc/pgx"
)

//Объект обращения к базе данных, параметр connSpec задан объектом Config пакета config
type DBProject struct {
	connSpec string
}

//NewDBProject() инициализирует DBProject, и возвращает указатель на него.
//Аргумент connSpec задаёт адрес для соединения с базой данных
func NewDBProject(connSpec string) *DBProject {
	return &DBProject{
		connSpec: connSpec,
	}
}

// Read() реализует чтение данных по конкретному проекту исходя из параметра логина projectId, заданных
// в структуре types.Project
func (d *DBProject) Read(projectId int64) (*types.Project, error) {
	return d.pgxConnProject(projectId)
}

//pgxConnProject() - функция принимает аргумент табличного идентификатора реализует внутри себя объект типа *types.Project обращается к базе данных
//в соответствии с заданным запросом, заполняет и возвращает объект и ошибку
func (d *DBProject) pgxConnProject(arg int64) (*types.Project, error) {
	//Объявляем пустую структуру types.Project а также переменные которые примут значение из базы данных
	var id int64
	var stringTasksSource string
	resultStruct := types.Project{}
	//Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), d.connSpec)
	if err != nil {
		return nil, fmt.Errorf("не удается подключиться к базе: %v", err)
	}
	defer conn.Close(context.Background())
	//Считываем необходимые параметры для всей заданной таблицы, проверяем ошибки
	err = conn.QueryRow(context.Background(), "SELECT proj.c_id, proj.c_name, CAST(proj.c_sprint_start_date AS TEXT), tref.c_name, proj.c_task_list_file_path FROM public.t_project AS proj LEFT JOIN t_ref_tasks_source AS tref ON proj.fk_tasks_source = tref.c_id WHERE proj.c_id=$1", arg).Scan(&id, &resultStruct.Name, &resultStruct.SprintStartDate, &stringTasksSource, &resultStruct.TaskListFilePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос к БД через метод pgxConnProject: %v", err)
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
