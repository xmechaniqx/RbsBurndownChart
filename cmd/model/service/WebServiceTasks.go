package service

import "net/http"

type WebServiceTasks struct {
	host   string
	port   string
	client *http.Client
}

func New(host string, port string) *WebServiceTasks {
	return nil
}
