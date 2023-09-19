package main

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/model"
	"RbsBurndownChart/cmd/types"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status        int8
	Error         string
	BurndownChart *types.Burndownchart
}

func main() {
	_, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
	outResponse := Response{}
	returner, err := model.MakeBurndownChart(login)
	if err != nil {
		outResponse = Response{
			Status:        1,
			Error:         "Ошибка! Неверный логин",
			BurndownChart: returner,
		}
	} else {
		outResponse = Response{
			Status:        0,
			Error:         "",
			BurndownChart: returner,
		}
	}
	//формируем JSON и отправляем на страницу "/chart?login..."
	output, err := json.MarshalIndent(outResponse, "", "\t")
	if err != nil {
		fmt.Println("Can't Marshall JSON", output)
	}
	w.Write(output)
}

// loadConfig() - функция загрузки файла конфигурации
func loadConfig() (*config.Config, error) {
	err := config.Load(parseParams())
	if err != nil {
		return nil, fmt.Errorf("error of load config: %s", err)
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
