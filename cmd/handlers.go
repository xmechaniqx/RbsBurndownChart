package main

import (
	"html/template"
	"log"
	"net/http"
)

//home() функция явлется отправной функцией для конфигурации ответа для клиентского запроса
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//Чтение файлов необходимых для отображения страницы HTML
	files := []string{
		"./ui/html/index.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
