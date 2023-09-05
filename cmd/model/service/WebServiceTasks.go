package service

import "net/http"

type WebServiceTasks struct {
	host   string
	port   string
	client *http.Client
}

//New() - функция получения списка задач из веб сервиса
func New(host string, port string) *WebServiceTasks {
	return nil
}
