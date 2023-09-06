package model

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/types"
)

//Объект обращения к пакету config.
type taskReaderFactory struct {
	config *config.Config
}

//New() инициализирует taskReaderFactory, и возвращает указатель на него.
//Аргумент config передает значение структуры Config пакета config.
func New(config *config.Config) *taskReaderFactory {
	return &taskReaderFactory{
		config: config,
	}
}

//Make() реализует имплементацию заданного типа проекта интерфейсу taskReader.
func (tr *taskReaderFactory) Make(project *types.Project) taskReader {

	return nil
}
