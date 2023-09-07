package main

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/model"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	runServer()

}

//responseHandler() - функция обработки ответа
func responseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// var tpl = template.Must(template.ParseFiles("index.html"))
	// tpl.Execute(w, nil)
	login := r.URL.Query().Get("login")

	fmt.Println(login)
	config, err := loadConfig()
	fmt.Println("full config?", *config)
	if err != nil {
		fmt.Errorf("ошибка чтения объекта \"Developer\" из базы данных: %v", err)
	}
	model := model.New(config)

	returner, err := model.MakeBurndownChart(login)
	if err != nil {
		fmt.Errorf("ошибка чтения объекта \"MakeBurndownChart\" из базы данных: %v", err)
	}
	fmt.Println(returner)
	output, err := json.MarshalIndent(returner, "", "\t")
	if err != nil {
		fmt.Println("Can't Marshall JSON")
	}
	w.Write(output)

}

//runServer() - функция запуска сервера
func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/chart", responseHandler)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	//Определение файлов необходимых для работы сервера
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

//loadConfig() - функция загрузки файла конфигурации
func loadConfig() (*config.Config, error) {
	err := config.Load(parseParams())
	if err != nil {
		fmt.Println("MAIN - Ошибка чтения файла config или пути configPath")
	}
	config := config.Read()
	return &config, err
}

//parseParams() - функция получения расположения файла conf.ini. Задается с помощью флага -p в терминале перед запуском приложения.
func parseParams() string {
	configPath := flag.String("p", "", "path of src")
	flag.Parse()
	return *configPath
}
