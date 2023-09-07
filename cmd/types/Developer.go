package types

// Объект содержащий данные о сотруднике, реализуется за счет соответствующей таблицы базы данных.
type Developer struct {
	Id       int64   //
	Projects []int64 //
	FullName string
	Login    string
	Position string
}
