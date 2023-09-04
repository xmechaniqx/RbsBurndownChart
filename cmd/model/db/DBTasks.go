package db

type DBTasks struct {
	connSpec string
	db       *pgx.Conn
}

func NewDBTasks(connSpec string) *DBTasks {
	return &DBTasks{}
}
