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
	loadConfig()
	runServer()

}

//runServer() - функция запуска сервера
func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/chart", chartHandler)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	//Определение файлов необходимых для работы сервера
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

//responseHandler() - функция обработки ответа
func chartHandler(w http.ResponseWriter, r *http.Request) {
	//Задаем тип заголовка для отображаемой страницы
	w.Header().Add("Content-Type", "application/json")
	//Вычленяем логин из адресной строки
	login := r.URL.Query().Get("login")
	//Зачитываем файл конфигурации
	config := config.Read()
	//Реализуем объект имплементирующий интерфейсу taskReader
	model := model.New(&config)
	//Получаем структуру с параметрами необходимыми для передачи в javaScript
	returner, err := model.MakeBurndownChart(login)
	if err != nil {
		fmt.Printf("ошибка чтения объекта \"MakeBurndownChart\" из базы данных: %v\n", err)
		return
	}
	//формируем JSON и отправляем на страницу "/chart?login..."
	output, err := json.MarshalIndent(returner, "", "\t")
	if err != nil {
		fmt.Println("Can't Marshall JSON", output)
	}
	w.Write(output)

}

// loadConfig() - функция загрузки файла конфигурации
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
	fmt.Println("parse")
	return *configPath
}
