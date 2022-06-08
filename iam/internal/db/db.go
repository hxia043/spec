package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DsnSuffix     string = "?_fk=true&_busy_timeout=30000"
	CreateUserSQL string = "create table if not exists User(UID integer not null primary key, Name text not null, Password text);"
)

func InitDatabase(path string) error {
	db, err := sql.Open("sqlite3", path+DsnSuffix)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec(CreateUserSQL); err != nil {
		return err
	}

	return nil
}

func OpenDatabase(path string) (*sql.DB, error) {
	dbConnection, err := sql.Open("sqlite3", path+DsnSuffix)
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}
