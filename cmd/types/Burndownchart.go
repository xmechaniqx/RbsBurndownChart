package types

// Объект данных содержащий данные о разработчике, и объекте Dataset из пакета types.
type Burndownchart struct {
	Developer Developer
	Charts    []Dataset
}
