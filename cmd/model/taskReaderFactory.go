package model

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/model/db"
	"RbsBurndownChart/cmd/model/file"
	"RbsBurndownChart/cmd/model/service"
	"RbsBurndownChart/cmd/types"
)

//Объект обращения к пакету config для получения структуры Config.
type taskReaderFactory struct {
	config *config.Config
}

//factory() инициализирует taskReaderFactory, и возвращает указатель на него.
//Аргумент config передает значение структуры Config пакета config.
func factory(config *config.Config) *taskReaderFactory {
	return &taskReaderFactory{
		config: config,
	}
}

//Make() реализует имплементацию заданного типа проекта интерфейсу taskReader в зависимости от источника задач.
//1. Для задач полученных из файла - types.File
//2. Для задач полученных из базы данных - types.DB
//3. Для задач полученных из веб-сервиса - types.Service
func (tr *taskReaderFactory) Make(project *types.Project) taskReader {
	switch {
	case project.TasksSource == types.File:
		return file.New("filepath from, db")
	case project.TasksSource == types.DB:
		return db.NewDBTasks(tr.config.DBConnSpec)
	case project.TasksSource == types.Service:
		return service.New(tr.config.ServiceHost, tr.config.ServicePort)
	}
	return nil
}
