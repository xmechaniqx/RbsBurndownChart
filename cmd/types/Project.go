package types

// Объект данных содержащий данные о проекте, реализуется за счет соответствующей таблицы базы данных.
type Project struct {
	Name             string
	SprintStartDate  string
	TasksSource      TaskSource
	TaskListFilePath string
}
