package service

import (
	"RbsBurndownChart/cmd/types"
	"net/http"
)

// Объект, реализующий соединение с веб-сервисом.
// Исходные данные для структуры хранятся в файле conf.ini далее в пакете Config.
type WebServiceTasks struct {
	host   string
	port   string
	client *http.Client
}

//New() - функция получения списка задач из веб-сервиса
func New(host string, port string) *WebServiceTasks {
	return nil
}

//Read() - функция чтения задач полученных из веб-сервиса
func (tr *WebServiceTasks) Read(p types.Project) ([]types.Task, error) {
	mockTasks := []types.Task{
		types.Task{
			Title: "Dummy task web 1",
			Cost:  3,
		},
		types.Task{
			Title: "Dummy task web 2",
			Cost:  1,
		},
		types.Task{
			Title: "Dummy task web 2",
			Cost:  1,
		},
	}
	return mockTasks, nil
}
