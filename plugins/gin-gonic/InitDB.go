package gingonic

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

func InitializeDB() *sql.DB{
	db,err := sql.Open("sqlite3","plugins/repositorys/sqliteRepository/db/db.db")
	if err != nil{
		panic(err)
	}

	_,err=db.Exec(`
	  CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		password TEXT NOT NULL 
	  )
	`)
	if err != nil{
		panic(err)
	}
	return db
}