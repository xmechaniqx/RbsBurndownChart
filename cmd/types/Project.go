package types

type Project struct {
	Name             string
	SprintStartDate  string
	TasksSource      TaskSource
	TaskListFilePath string
}
