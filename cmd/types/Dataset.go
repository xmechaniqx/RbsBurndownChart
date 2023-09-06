package types

// Объект данных содержащий данные о проекте, сумму времени (в часах) на выполнение проекта,
// и отображение расчетных часов работы в соответствии с днями недели.
type Dataset struct {
	Project      Project
	TasksCostSum int64
	WorkingHours map[string]int64
}
