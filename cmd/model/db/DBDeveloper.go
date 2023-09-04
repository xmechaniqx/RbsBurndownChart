package db

type DBDeveloper struct {
	connSpec string
	db       *pgx.Conn
}

func NewDBDeveloper(connSpec string) *DBDeveloper {
	return &DBDeveloper{}
}

func (db *DBDeveloper) Read(devLogin string) (*types.Developer, error) {
	return nil, nil
}
func ReadAll(projectId int64) ([]types.Developer, error) {
	return nil, nil
}
