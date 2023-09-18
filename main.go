package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	runServer()

}

type Task struct {
	Title string //Название задачи
	Cost  int64  //Стоимость задачи
}

//runServer() - функция запуска сервера
func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// mux.HandleFunc("/chart", chartHandler)
	log.Println("Запуск TASK веб-сервера на http://127.0.0.1:5000")
	//Определение файлов необходимых для работы сервера
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	returner, err := readLines("/home/username/go/src/RbsBurndownChart/tasks/group_3.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла веб-сервер")
	}
	output, err := json.MarshalIndent(returner, "", "\t")
	if err != nil {
		fmt.Println("Can't Marshall JSON", output)
	}
	w.Write(output)

}
func readLines(path string) ([]Task, error) {
	//Инициализируем переменные которые будут записаны в структуру
	var line string
	var resInt int
	//Инициализируем структуру и массив структур
	arrResult := []Task{}
	result := Task{}
	//Чтение файла
	file, err := os.Open(path)
	if err != nil {
		return []Task{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		//Форматируем полученную строку чтобы вычленить время и удалить лишние символы
		rex := regexp.MustCompile(`\(([^)]+)\)`)
		out := rex.FindAllStringSubmatch(line, -1)
		line = strings.TrimRight(line, "()+1234567890")
		line = strings.TrimLeft(line, ".()+1234567890")
		fmt.Println("rex", rex, "out", out, "line", line)
		//Приводим полученное число из строки в int
		for _, i := range out {
			resInt, err = strconv.Atoi(i[1])
			if err != nil {
				fmt.Println("Ошибка конвертации STRING to INT", err)
			}
		}
		//Записываем полученные значения в структуру
		result = Task{
			Title: line,
			Cost:  int64(resInt),
		}
		//Записываем полученную структуру в массив
		arrResult = append(arrResult, result)
	}
	return arrResult, err
}

// func main() {
// 	res, err := readLines("/home/username/go/src/RbsBurndownChart/tasks/group_1.txt")
// 	if err != nil {
// 		return
// 	}
// 	fmt.Printf("%+v\n", res)
// }
// func readLines(path string) ([]Task, error) {
// 	var line string
// 	var resInt int
// 	arrResult := []Task{}
// 	result := Task{}
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return []Task{}, err
// 	}
// 	defer file.Close()

// 	// var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line = scanner.Text()
// 		rex := regexp.MustCompile(`\(([^)]+)\)`)
// 		out := rex.FindAllStringSubmatch(line, -1)
// 		line = strings.TrimRight(line, "()+1234567890")
// 		line = strings.TrimLeft(line, ".()+1234567890")
// 		for _, i := range out {
// 			resInt, err = strconv.Atoi(i[1])
// 			if err != nil {
// 				fmt.Println("Ошибка конвертации STRING to INT")
// 			}
// 		}

// 		result = Task{
// 			Title: line,
// 			Cost:  int64(resInt),
// 		}
// 		arrResult = append(arrResult, result)
// 	}

// 	// fmt.Println(sumTaskHoursArr, sumResult)
// 	return arrResult, err
// }

// type Task struct {
// 	Title string
// 	Cost  int64
// }
