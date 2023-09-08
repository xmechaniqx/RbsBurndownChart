package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	res, err := readLines("/home/username/go/src/RbsBurndownChart/tasks/group_1.txt")
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", res)
}
func readLines(path string) ([]Task, error) {
	var line string
	var resInt int
	arrResult := []Task{}
	result := Task{}
	file, err := os.Open(path)
	if err != nil {
		return []Task{}, err
	}
	defer file.Close()

	// var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		rex := regexp.MustCompile(`\(([^)]+)\)`)
		out := rex.FindAllStringSubmatch(line, -1)
		line = strings.TrimRight(line, "()+1234567890")
		line = strings.TrimLeft(line, ".()+1234567890")
		for _, i := range out {
			resInt, err = strconv.Atoi(i[1])
			if err != nil {
				fmt.Println("Ошибка конвертации STRING to INT")
			}
		}

		result = Task{
			Title: line,
			Cost:  int64(resInt),
		}
		arrResult = append(arrResult, result)
	}

	// fmt.Println(sumTaskHoursArr, sumResult)
	return arrResult, err
}

type Task struct {
	Title string
	Cost  int64
}
