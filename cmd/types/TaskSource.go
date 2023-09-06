package types

// Объект данных содержащий данные о способе получения задач для проекта.
type TaskSource int8

const (
	File    TaskSource = 1
	DB      TaskSource = 2
	Service TaskSource = 3
)
