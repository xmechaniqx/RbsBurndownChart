package main

import (
	"RbsBurndownChart/cmd/config"
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	loadConfig()
	runServer()

}

//responseHandler() - функция обработки ответа
func responseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// var tpl = template.Must(template.ParseFiles("index.html"))
	// tpl.Execute(w, nil)
	login := r.URL.Query().Get("login")
	fmt.Println(login)
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
	err := config.Load(*parseParams())
	if err != nil {
		fmt.Println("MAIN - Ошибка чтения файла config или пути configPath")
	}
	myConfig := config.Read()
	return &myConfig, err
}

//parseParams() - функция получения расположения файла conf.ini. Задается с помощью флага -p в терминале перед запуском приложения.
func parseParams() (configPath *string) {
	configPath = flag.String("p", "", "path of src")
	flag.Parse()
	return configPath
}

///////////////////////////////////////////////////////

//Построчное считывание файла задач
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
	// path := "./tasks/"
	// var curPath string
	// var sumResult int
	// files, err := ioutil.ReadDir(path)
	// fmt.Println(files)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, file := range files {
	// 	curPath = path + file.Name()
	// 	sumResult = takePathGiveSum(curPath)
	// 	fmt.Println("This is result of sum for", file.Name(), "-", sumResult)
	// }
}

//Функция вычисления суммы для экземпляра задачи
func takePathGiveSum(path string) int {
	var taskString []string
	var sumResult int
	taskString, err := readLines(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
	}
	for _, v := range taskString {
		rex := regexp.MustCompile(`\(([^)]+)\)`)
		out := rex.FindAllStringSubmatch(v, -1)
		for _, i := range out {
			resInt, err := strconv.Atoi(i[1])
			if err != nil {
				fmt.Println("Ошибка конвертации STRING to INT")
			}
			sumResult += resInt
		}
	}
	// fmt.Println(sumTaskHoursArr, sumResult)
	return sumResult
}
