package utils

import (
	"database/sql"
	"sync"
)

type Document struct {
	mu sync.Mutex
	db *sql.DB
}
