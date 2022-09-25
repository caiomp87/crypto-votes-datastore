package postgres

import "database/sql"

type IPostres interface {
	Connect() (*sql.DB, error)
	Disconnect(*sql.DB) error
}

type PostgresService struct {
	driver           string
	connectionString string
}

func NewPostresService(driver, connectionString string) IPostres {
	return &PostgresService{
		driver,
		connectionString,
	}
}

func (ps *PostgresService) Connect() (*sql.DB, error) {
	db, err := sql.Open(ps.driver, ps.connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (ps *PostgresService) Disconnect(db *sql.DB) error {
	return db.Close()
}
