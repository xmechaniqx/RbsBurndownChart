package model

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/model/db"
	"RbsBurndownChart/cmd/types"
	"fmt"
)

// Объект, реализующий чтение из базы данных параметров имплементирующих интерфейсу taskReader.
type Provider struct {
	taskReadeFactory *taskReaderFactory
	projects         *db.DBProject
	devs             *db.DBDeveloper
}

//NewProvider() - функция обработчик структуры Config, реализующая доступ к базе данных согласно файлу config.ini
func New(config *config.Config) *Provider {
	return &Provider{
		taskReadeFactory: factory(config),
		devs:             db.NewDBDeveloper(config.DBConnSpec),
		projects:         db.NewDBProject(config.DBConnSpec),
	}
}

//MakeBurndownChart() - функция получения структуры данных необходимых для построения графика согласно логину пользователя
func (p *Provider) MakeBurndownChart(devLogin string) (*types.Burndownchart, error) {
	var dataSetsDeveloper []types.Dataset

	//Чтение данных сотрудника из БД
	developerObject, err := p.devs.Read(devLogin)
	// fmt.Println("developerObject.MakeBurndownChart", developerObject)
	if err != nil || developerObject == nil {
		return nil, fmt.Errorf("ошибка чтения объекта \"Developer\" из базы данных: %v", err)
	}

	//Заполняем DataSet в зависимости от количества связанных с сотрудником проектов
	for i := 0; i < len(developerObject.Projects); i++ {
		dataSet, err := p.makeDataSet(developerObject.Projects[i])

		if err != nil {
			return nil, fmt.Errorf("ошибка сборки объекта \"dataSet\": %v", err)
		}
		dataSetsDeveloper = append(dataSetsDeveloper, *dataSet)
	}

	//Инициализация структуры BurndownChart и запись полученных значений
	burnDown := &types.Burndownchart{
		Developer: *developerObject,
		Charts:    dataSetsDeveloper,
	}
	return burnDown, err
}

//makeDataSet() - функция принимает параметр ID проекта, обращается к базе данных и возвращает объект
func (p *Provider) makeDataSet(projectId int64) (*types.Dataset, error) {

	//Чтение проекта из БД
	project, err := p.projects.Read(projectId)
	// fmt.Println("project is", project)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения объекта \"Project\" из базы данных: %v", err)
	}

	//Инициализация taskReader через factory()
	dataSetTaskReaderMake := p.taskReadeFactory.Make(project)
	dataSetTaskReader, err := dataSetTaskReaderMake.Read(*project)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения задачи : %v", err)
	}

	//Вычисление суммы стоимостей задач
	sum := p.costSum(dataSetTaskReader)

	//Чтение всех сотрудников проекта
	allDevsOfProject, err := p.readProjectTeam(projectId)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения всех объектов \"Developer\" из базы данных : %v", err)
	}

	//Заполнение рабочих часов сотрудников
	workingHoursOfAll, err := p.calkWorkingHours(allDevsOfProject)

	//Инициализация структуры Dataset и запись полученных значений
	chart := &types.Dataset{
		Project:      *project,
		TasksCostSum: sum,
		WorkingHours: workingHoursOfAll,
	}
	return chart, err
}

//costSum() - функция возвращающая суммарное время необходимое для решения задачи
func (p *Provider) costSum(tasks []types.Task) int64 {
	var myCostSum int64
	for _, task := range tasks {
		myCostSum += task.Cost
	}
	return myCostSum
}

//readProjectTeam() - функция возвращает список участников для заданного проекта
func (p *Provider) readProjectTeam(projectId int64) ([]types.Developer, error) {
	myTeamAll, err := p.devs.ReadAll(projectId)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения всех объектов \"Developer\" из базы данных : %v", err)
	}
	return myTeamAll, err
}

//readWorkingHours() - функция возвращает производительность работника в соответсвии с недельным графиком
func (p *Provider) calkWorkingHours(devs []types.Developer) (map[string]int64, error) {
	sumDayToHours := make(map[string]int64)
	// for _, developer := range devs {
	// 	someDevs:=db.NewDBDeveloper()

	sumDayToHours["mon"] += 8
	sumDayToHours["tue"] += 8
	sumDayToHours["wed"] += 6
	sumDayToHours["thu"] += 3
	sumDayToHours["fri"] += 7
	sumDayToHours["sat"] += 0
	sumDayToHours["sun"] += 0
	// }
	// sumDayToHours["mon"]+=db.DBDeveloper...SELECT FROM c_mon_hour
	// sumDayToHours["tue"]+=db.DBDeveloper...SELECT FROM c_tue_hour
	// sumDayToHours["wed"]+=db.DBDeveloper...SELECT FROM c_wed_hour
	// sumDayToHours["thu"]+=db.DBDeveloper...SELECT FROM c_thu_hour
	// sumDayToHours["fri"]+=db.DBDeveloper...SELECT FROM c_fri_hour
	// sumDayToHours["sat"]+=db.DBDeveloper...SELECT FROM c_sat_hour
	// sumDayToHours["sun"]+=db.DBDeveloper...SELECT FROM c_sun_hour
	return sumDayToHours, nil
}

//Интерфейс определяющий задачи с каким расположением будут переданы для реализации в пакете Provider
type taskReader interface {
	Read(p types.Project) ([]types.Task, error)
}
