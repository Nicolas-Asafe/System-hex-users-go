package main

import (
	"database/sql"
	"fmt"
	"log"

	// "pc/core/models"
	"pc/core/services"
	"pc/plugins/repositorys/sqliteRepository"

	_ "github.com/mattn/go-sqlite3"
)

var serv services.UserService // Variável global para uso em outras funções

func main() {
	// Conectar ao banco
	db, err := sql.Open("sqlite3", "./plugins/repositorys/sqliteRepository/db/db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela se não existir
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	repo := sqliterepository.SqliteRepoInitialize(db)
	serv = services.UserService{Repo: repo}

	// Criar um novo usuário
	// serv.NewUser(models.User{
	// 	Name:     "Nicolas",
	// 	Password: "1234",
	// })

	// Deletar um usuário com id 1
	err = serv.DeleteUserById(2)
	if err != nil {
		fmt.Println("Erro ao deletar:", err)
	} else {
		fmt.Println("Usuário deletado com sucesso!")
	}

	// Listar os usuários
	PrintUsers()
}

func PrintUsers() {
	users, err := serv.ListUsers()
	if err != nil {
		log.Fatal(err)
	}

	if len(users) == 0 {
		fmt.Println("Nenhum usuário encontrado.")
		return
	}

	fmt.Println("Lista de usuários:")
	fmt.Println("---------------------------")
	for _, user := range users {
		fmt.Printf("ID: %-3d | Nome: %-20s\n", user.Id, user.Name)
		fmt.Println("---------------------------")
	}
}
