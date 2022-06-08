package models

import (
	"database/sql"
	"sync"
)

var DbConnection *sql.DB
var DbRwMutex sync.RWMutex
