package main

import (
	"RbsBurndownChart/cmd/config"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	path := "./tasks/"
	var curPath string
	var sumResult int
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		curPath = path + file.Name()
		sumResult = takePathGiveSum(curPath)
		fmt.Println("This is result of sum for", file.Name(), "-", sumResult)
	}

}

//responseHandler() - функция обработки ответа
func responseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	root := r.URL.Query().Get("chart")
	// flag?root=/home/username/node_modules/
	defaulPath := "/home/username/node_modules/"
	if root == "/" {
		root = defaulPath
		fmt.Println("defaulRoot")
	}
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

//parseParams() - функция обработки файла конфигурации
func parseParams() configPath {
	var i configPath
	return i
}

//loadConfig() - функция загрузки файла конфигурации
func loadConfig() (*config.Config, error) {
	return nil, nil
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
