package types

// Объект данных содержащий данные о способе получения задач для проекта.
type TaskSource int8

const (
	File    TaskSource = 1 //Задачи расположены в файле
	DB      TaskSource = 2 //Задачи расположены в БД
	Service TaskSource = 3 //Задачи расположены на веб-сервисе
)
