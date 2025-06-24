package main

import (
	"database/sql"
	"pc/core/models"
	"pc/core/services"
	"pc/plugins/repositorys/sqliteRepository"
)

func main() {
	db,err:= sql.Open("sqlite3","./plugins/repositorys/sqliteRepository/db/db.db")
	if err != nil{
		panic(err)
	}
	defer db.Close()
	_,err= db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 password TEXT NOT NULL
	)
	`)
	if err != nil{
		panic(err)
	}
	repo := sqliterepository.SqliteRepoInitialize(db)
	serv := services.UserService{Repo: repo}
	serv.NewUser(models.User{
		Name: "Nicolas",
		Password: "1234",
	})
	users,err := serv.ListUsers()
	if err != nil{
		panic(err)
	}
	for _,user := range users{
		println(user.Name)
	}
}