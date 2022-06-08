package models

import (
	"iam/internal/db"
	"os"
	"path/filepath"
)

const (
	InsertUserSQL = "insert into User(UID, Name, Password) values (?,?,?);"
)

func InsertUser(name, password string) error {
	var err error = nil
	dbPath := filepath.Join(os.Getenv("HOME"), "sqlite3_iam.db")

	DbConnection, err = db.OpenDatabase(dbPath)
	if err != nil {
		return err
	}
	defer DbConnection.Close()

	DbRwMutex.Lock()
	defer DbRwMutex.Unlock()

	if _, err := DbConnection.Exec(InsertUserSQL, 100, name, password); err != nil {
		return err
	}

	return nil
}
