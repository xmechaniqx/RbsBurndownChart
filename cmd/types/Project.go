package types

// Объект данных содержащий данные о проекте, реализуется за счет соответствующей таблицы базы данных.
type Project struct {
	Id               int64      //Идентификатор проекта
	Name             string     //Название проекта
	SprintStartDate  string     //Дата начала спринта для проекта
	TasksSource      TaskSource //Параметр типа расположения списка задач
	TaskListFilePath string     //Расположение списка задач в виде файла на сервере
}
