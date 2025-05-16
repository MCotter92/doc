package utils

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	mu sync.Mutex
	db *sql.DB
}

const create = `
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY, -- UUID as TEXT
    name TEXT NOT NULL,
    notes_location TEXT NOT NULL
, editor text NOT NULL);
`

const file string = "database.db"

func CreateUserTable() (*database, error) {
	//create the user table in database.db

	// init db handler
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Printf("cannot init %v handle: %v\n", db, err)
	}

	//create user table if it doesn't exists yet
	sqlResult, err := db.Exec(create)
	if err != nil {
		fmt.Println("Couldn't create user DB: ", err)
	}

	fmt.Println(sqlResult)
	return &database{db: db}, nil

}

func (u *database) CreateUser(name, notesLocation, editor string) (string, error) {
	// insert a new user into the users table in database.db

	id := uuid.New().String()

	_, err := u.db.Exec(`
		INSERT INTO users (id, name, notes_location, editor)
		VALUES (?, ?, ?, ?)
	`, id, name, notesLocation, editor)

	if err != nil {
		fmt.Println("Cannot insert into DB:", err)
		return "", err
	}

	return id, nil
}

// func (u *User) CreateUser(name, notes_location, editor string) (int, error) {
// 	res, err := u.db.Exec("INSERT INTO Users(id, name, notes_location, editor) VALUES(????);", uuid.New.string(), name, notes_location, editor)
// 	if err != nil {
// 		fmt.Println("Cannot Insert into DB.", err)
// 	}
//
// 	var id int64
// 	if id, err = res.LastInsertId(); err != nil {
// 		return 0, err
// 	}
// 	return int(id), err
//
// }
