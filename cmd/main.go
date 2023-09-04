package main

import (
	"RbsBurndownChart/cmd/config"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
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
func runServer() {

}
func parseParams() configPath {
	var i configPath
	return i
}

func loadConfig() (*config.Config, error) {
	return nil, nil
}

///////////////////////////////////////////////////////

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
