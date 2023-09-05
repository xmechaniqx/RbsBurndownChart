package model

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/model/db"
	"RbsBurndownChart/cmd/types"
)

type Provider struct {
	tasks       taskReader
	projects    *db.DBProject
	devs        *db.DBDeveloper
	currentUser *types.Developer
}

//NewProvider() - функция обработчик структуры Config, реализующая доступ к базе данных согласно файлу config.ini
func NewProvider(config *config.Config) *Provider {
	return nil
}

//MakeBurndownChart() - функция получения структуры данных необходимых для построения графика согласно логину пользователя
func (p *Provider) MakeBurndownChart(devLogin string) (*types.Burndownchart, error) {
	return nil, nil
}

//readUser() - функция чтения из базы данных информации о пользователе
func (p *Provider) readUser(devLogin string) (*types.Developer, error) {
	return nil, nil
}

//readProject() - функция чтения из базы данных информации о проекте
func (p *Provider) readProject(projectId int64) (*types.Project, error) {
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

type taskReader interface {
	Read(p *types.Project) ([]types.Task, error)
}
