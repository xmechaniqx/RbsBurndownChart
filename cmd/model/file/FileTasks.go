package file

import (
	"RbsBurndownChart/cmd/types"
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
	mockTasks := []types.Task{
		types.Task{
			Title: "Dummy task 1",
			Cost:  8,
		},
		types.Task{
			Title: "Dummy task 2",
			Cost:  3,
		},
	}
	return mockTasks, nil
}
