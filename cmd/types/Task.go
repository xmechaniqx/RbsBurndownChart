package types

// Объект данных содержащий данные о стоимости конкретной задачи в проекте, реализуется за счет соответствующей таблицы базы данных.
type Task struct {
	Title string //Название задачи
	Cost  int64  //Стоимость задачи
}
