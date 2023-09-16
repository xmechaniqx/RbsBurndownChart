package db

import (
	"RbsBurndownChart/cmd/types"
	"context"
	"fmt"
	"os"

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
	return pgxConnDev(devLogin)
}

// ReadAll() реализует чтение списка разработчиков для проекта (projectId) из базы данных, заданной
// в структуре DBDeveloper.
func (db *DBDeveloper) ReadAll(projectId int64) ([]types.Developer, error) {
	return pgxConnAllDevs(projectId)
}
func pgxConnDev(devLogin string) (*types.Developer, error) {
	//Объявляем пустую структуру types.Project а также переменные которые примут значение из базы данных
	var id, projectID int64
	var projectIDs []int64
	var lastname, firstname, position string
	resultStruct := types.Developer{}
	resultMapHours := make(map[string]int64)
	var mon, tue, wed, thu, fri, sat, sun int64
	//Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://admin:1234@localhost:5432/burndown_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	//Считываем необходимые параметры для всей заданной таблицы, проверяем ошибки
	all, err := conn.Query(context.Background(), "SELECT dev.c_id, toc.fk_project, dev.c_lastname, dev.c_firstname, tref.c_name, dev.c_mon_hour, dev.c_tue_hour, dev.c_wed_hour,	dev.c_thu_hour, dev.c_fri_hour, dev.c_sat_hour, dev.c_sun_hour FROM t_developer AS dev LEFT JOIN t_ref_position AS tref ON dev.fk_position = tref.c_id	LEFT JOIN toc_project_developer AS toc ON dev.c_id = toc.fk_developer WHERE dev.c_login =$1", devLogin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow for one Developer failed: %v\n", err)
		os.Exit(1)
	}
	//Цикл чтения строк таблицы и запись данных в структуру taskStruct, далее заполненяем массив объектов полученной структурой
	for all.Next() {
		err := all.Scan(&id, &projectID, &lastname, &firstname, &position, &mon, &tue, &wed, &thu, &fri, &sat, &sun)
		if err != nil {
			fmt.Println("next error")
		}
		if all.Err() != nil {
			fmt.Println("scan error")
		}
		resultMapHours = map[string]int64{
			"Понедельник": mon,
			"Вторник":     tue,
			"Среда":       thu,
			"Четверг":     wed,
			"Пятница":     fri,
			"Суббота":     sat,
			"Воскресенье": sun,
		}
		projectIDs = append(projectIDs, projectID)
	}
	if lastname == "" || firstname == "" {
		fmt.Println("Have no developer")
		return nil, err
	}
	// fmt.Println("eto karta", resultMapHours)
	//Заполняем структуру полученными значениями
	resultStruct = types.Developer{
		Id:                id,
		Projects:          projectIDs,
		FullName:          lastname + " " + firstname,
		Login:             devLogin,
		Position:          position,
		WorkingHoursOfOne: resultMapHours,
	}
	return &resultStruct, err
}
func pgxConnAllDevs(projectId int64) ([]types.Developer, error) {
	//Объявляем пустую структуру types.Project а также переменные которые примут значение из базы данных
	var allProjectIDs []types.Developer
	var login string
	//Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://admin:1234@localhost:5432/burndown_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	//Считываем необходимые параметры для всей заданной таблицы, проверяем ошибки
	allLogin, err := conn.Query(context.Background(), "SELECT dev.c_login FROM t_developer AS dev LEFT JOIN t_ref_position AS tref ON dev.fk_position = tref.c_id LEFT JOIN toc_project_developer AS toc	ON dev.c_id = toc.fk_developer WHERE toc.fk_project=$1", projectId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow for all developers is failed: %v\n", err)
		os.Exit(1)
	}
	//Цикл чтения строк таблицы и запись данных в структуру taskStruct, далее заполненяем массив объектов полученной структурой
	for allLogin.Next() {
		err := allLogin.Scan(&login)
		if err != nil {
			fmt.Println("next error")
		}
		if allLogin.Err() != nil {
			fmt.Println("scan error")
		}
		oneDev, err := pgxConnDev(login)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow for all developers is failed: %v\n", err)
			os.Exit(1)
		}
		//Заполняем массив объектами types.Developer
		allProjectIDs = append(allProjectIDs, *oneDev)
	}
	return allProjectIDs, err
}
