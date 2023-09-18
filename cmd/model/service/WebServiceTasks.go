package service

import (
	"RbsBurndownChart/cmd/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
	var myClient = &http.Client{Timeout: 10 * time.Second}
	returnWebServiceTasks := WebServiceTasks{
		host:   host,
		port:   port,
		client: myClient,
	}
	return &returnWebServiceTasks
}

//Read() - функция чтения задач полученных из веб-сервиса
func (tr *WebServiceTasks) Read(p types.Project) ([]types.Task, error) {
	//Инициализируем новый массив объектов типа types.Task
	mockTasks := []types.Task{}
	//Совершаем запрос на сервер в соответсвии с данными из структуры WebServiceTask, которая в свою очередь сконфигурирована из файла config
	r, err := tr.client.Get(tr.host + tr.port)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	//Приводим тип полученного ответа к байтовому срезу
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	//Декодируем полученные значения и заполняет массив структур
	if err := json.Unmarshal(body, &mockTasks); err != nil {
		fmt.Println(err)
	}
	return mockTasks, err
}
