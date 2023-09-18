package file

import (
	"RbsBurndownChart/cmd/types"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Объект, реализующий чтение списка задач из текстового файла.
// Путь к файлу с задачами задаётся параметром filePath конструктора.
// FileTasks имплементирует интерфейс tasksReader.
type FileTasks struct {
	filePath string // Путь к файлу со списком задач.
}

// New() инициализирует FileTasks, и возвращает указатель на него.
// filePath задаёт путь к файлу со списком задач.
func New(filePath string) *FileTasks {
	return &FileTasks{filePath: filePath}
}

// Read() реализует чтение списка задач из текстового файла, заданного
// в структуре FileTasks.
// Формат текстового файла:
// задача 1 (4)
// задача 2 (6)
// ...
// задача N (X)
// т.е. каждая новая задача на отдельной строке, а стоимость задачи указывается целым числом в круглых скобках.
func (tr *FileTasks) Read(p types.Project) ([]types.Task, error) {
	return readLines(p.TaskListFilePath)
}

//Функция построчного считывания и форматирования согласно объекту .txt файла задач с последующим формированием массива по каждой строке
func readLines(path string) ([]types.Task, error) {
	//Инициализируем переменные которые будут записаны в структуру
	var line string
	var resInt int
	//Инициализируем структуру и массив структур
	arrResult := []types.Task{}
	result := types.Task{}
	//Чтение файла
	file, err := os.Open(path)
	if err != nil {
		return []types.Task{}, err
	}
	defer file.Close()
	//Построчное сканирование файла
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		//Форматируем полученную строку чтобы вычленить время и удалить лишние символы
		rex := regexp.MustCompile(`\(([^)]+)\)`)
		out := rex.FindAllStringSubmatch(line, -1)
		line = strings.TrimRight(line, "()+1234567890")
		line = strings.TrimLeft(line, ".()+1234567890")
		//Приводим полученное число из строки в int
		for _, i := range out {
			resInt, err = strconv.Atoi(i[1])
			if err != nil {
				fmt.Println("Ошибка конвертации STRING to INT", err)
			}
		}
		//Записываем полученные значения в структуру
		result = types.Task{
			Title: line,
			Cost:  int64(resInt),
		}
		//Записываем полученную структуру в массив
		arrResult = append(arrResult, result)
	}
	return arrResult, err
}
