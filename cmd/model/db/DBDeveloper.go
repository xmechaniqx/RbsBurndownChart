package db

import (
	"RbsBurndownChart/cmd/types"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

//Объект обращения к базе данных, параметр connSpec задан объектом Config пакета config
type DBDeveloper struct {
	connSpec string
	db       *pgx.Conn
}

// NewDBDeveloper() инициализирует DBDeveloper, и возвращает указатель на него.
// Аргумент connSpec задаёт адрес для соединения с базой данных
func NewDBDeveloper(connSpec string) *DBDeveloper {
	return &DBDeveloper{}
}

// Read() реализует чтение данных по конкретному разработчику исходя из параметра логина devLogin, заданных
// в структуре types.Developer
func (d *DBDeveloper) Read(devLogin string) (*types.Developer, error) {
	return pgxConnDev(devLogin)
}

// ReadAll() реализует чтение списка разработчиков для проекта (projectId) из базы данных, заданной
// в структуре DBDeveloper
func (db *DBDeveloper) ReadAll(projectId int64) ([]types.Developer, error) {
	return pgxConnAllDevs(projectId)
}

//pgxConnDev() - функция принимает логин пользователя и обращается к БД для заполнения объекта types.Developer
func pgxConnDev(devLogin string) (*types.Developer, error) {
	//Объявляем пустую структуру types.Developer а также переменные, которые примут значение из базы данных
	resultStruct := types.Developer{}
	var id, projectID int64 // id сотрудника и id проекта
	var projectIDs []int64  // массив идентификаторов проектов для одного сотрудника, у сотрудника может быть больше одного проекта в работе
	var lastname, firstname, position string
	var mon, tue, wed, thu, fri, sat, sun int64
	resultMapHours := make(map[string]int64)
	//Подключаемся к базе данных
	// TODO:
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
	//Цикл чтения строк таблицы и запись данных в структуру resultStruct
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
		//собираем массив идентификаторов проекта если их больше чем один
		projectIDs = append(projectIDs, projectID)
	}
	if lastname == "" || firstname == "" {
		return nil, fmt.Errorf("have no developer %s", devLogin)
	}
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

//pgxConnAllDevs() - функция принимает идентификатор проекта и обращается к БД для заполнения массива объектов []types.Developer.
func pgxConnAllDevs(projectId int64) ([]types.Developer, error) {
	//Объявляем пустую структуру types.Developer а также переменные которые примут значение из базы данных
	var allProjectIDs []types.Developer
	var login string
	//Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://admin:1234@localhost:5432/burndown_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	//Считываем параметр логины всех пользователей которые рабтают над проектом с заданным идентификатором projectId, проверяем ошибки
	allLogin, err := conn.Query(context.Background(), "SELECT dev.c_login FROM t_developer AS dev LEFT JOIN t_ref_position AS tref ON dev.fk_position = tref.c_id LEFT JOIN toc_project_developer AS toc	ON dev.c_id = toc.fk_developer WHERE toc.fk_project=$1", projectId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow for all developers is failed: %v\n", err)
		os.Exit(1)
	}
	//Цикл чтения строк таблицы и вызов функции pgxConnDev() для получения структуры types.Developer для каждого пользователя и запись объектов в массив
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
