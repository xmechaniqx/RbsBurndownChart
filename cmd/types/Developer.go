package types

// Объект содержащий данные о сотруднике, реализуется за счет соответствующей таблицы базы данных.
type Developer struct {
	Id                int64            //Идентификатор сотрудника
	Projects          []int64          //Список проекта(или проектов) сотрудника
	FullName          string           //Полное имя сотрудника
	Login             string           //Логин
	Position          string           //Должность
	WorkingHoursOfOne map[string]int64 //Занятость сотрудника в часах по дням недели
}
