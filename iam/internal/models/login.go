package models

import (
	"fmt"
	"iam/internal/db"
	"iam/internal/pkg/hash"
	"os"
	"path/filepath"
)

const (
	QueryUserSQL = "select Password from User where Name = ?"
)

func Login(name, password string) error {
	var err error = nil
	dbPath := filepath.Join(os.Getenv("HOME"), "sqlite3_iam.db")

	DbConnection, err = db.OpenDatabase(dbPath)
	if err != nil {
		return err
	}
	defer DbConnection.Close()

	DbRwMutex.Lock()
	defer DbRwMutex.Unlock()

	var DbPassword string
	err = DbConnection.QueryRow(QueryUserSQL, name).Scan(&DbPassword)
	if err != nil {
		return err
	}

	if !hash.CheckPasswordHash(password, DbPassword) {
		return fmt.Errorf("error: password %s is not correct", password)
	}

	return nil
}
