package model

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/model/db"
	"RbsBurndownChart/cmd/types"
	"fmt"
)

// Объект, реализующий чтение из базы данных параметров имплементирующих интерфейсу taskReader.
type Provider struct {
	tasks       taskReader
	projects    *db.DBProject
	devs        *db.DBDeveloper
	currentUser *types.Developer
	config      *config.Config
}

//NewProvider() - функция обработчик структуры Config, реализующая доступ к базе данных согласно файлу config.ini
func New(config *config.Config) *Provider {
	return &Provider{
		config: config,
	}
}

//MakeBurndownChart() - функция получения структуры данных необходимых для построения графика согласно логину пользователя
func (p *Provider) MakeBurndownChart(devLogin string) (*types.Burndownchart, error) {

	// myDataset:=types.Dataset{
	// 	Project: developerObject.,
	// }
	return nil, nil
}

//readUser() - функция чтения из базы данных информации о пользователе
func (p *Provider) readUser(devLogin string) (*types.Developer, error) {
	developerMake := db.DBDeveloper(*db.NewDBDeveloper(p.config.DBConnSpec))
	developerObject, err := developerMake.Read(devLogin)
	if err != nil {
		fmt.Println(`Ошибка чтения параметра "Developer" из базы данных`)
	}
	return developerObject, err
}

//readProject() - функция чтения из базы данных информации о проекте
func (p *Provider) readProject(projectId int64) (*types.Project, error) {
	developerMake := db.DBProject(*db.NewDBProject(p.config.DBConnSpec))
	developerObject, err := developerMake.Read(devLogin)
	if err != nil {
		fmt.Println(`Ошибка чтения параметра "Developer" из базы данных`)
	}
	return developerObject, err
	return nil, nil
}

//readTasks() - функция чтения из базы данных списка задач
func (p *Provider) readTasks(projectId int64) ([]types.Task, error) {
	return nil, nil
}

//costSum() - функция возвращающая суммарное время необходимое для решения задачи
func (p *Provider) costSum(tasks []types.Task) (int64, error) {
	return 0, nil
}

//readProjectTeam() - функция возвращает список участников для заданного проекта
func (p *Provider) readProjectTeam(projectId int64) ([]types.Developer, error) {
	return nil, nil
}

//readWorkingHours() - функция возвращает производительность работника в соответсвии с недельным графиком
func (p *Provider) readWorkingHours(devs []types.Developer) (map[string]int64, error) {
	return nil, nil
}

//Интерфейс определяющий задачи с каким расположением будут переданы для реализации в пакете Provider
type taskReader interface {
	Read(p types.Project) ([]types.Task, error)
}
