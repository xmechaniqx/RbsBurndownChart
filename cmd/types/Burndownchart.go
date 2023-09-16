package types

// Объект данных содержащий данные о разработчике, и объекте Dataset из пакета types
type Burndownchart struct {
	Developer Developer //Объект сотрудник
	Charts    []Dataset //Набор данных о проекте(или проектах) в которых учавствует сотрудник
}
