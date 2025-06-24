package sqliterepository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"pc/core/models"
	"pc/core/interfaces"
)

type SqliteRepo struct {
	db *sql.DB
}

func (s *SqliteRepo) Save(user models.User) error {
	_, err := s.db.Exec(`INSERT INTO users(name, password) VALUES (?, ?)`, user.Name, user.Password)
	return err
}
func  (s *SqliteRepo) Find() ([]models.User,error){
    rows,err := s.db.Query(`SELECT name,password FROM users`)
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Name,&user.Password)
		if err != nil{
		   return nil,err
		}
		users = append(users, user)
	}
	return users,err
}

func SqliteRepoInitialize(db *sql.DB) interfaces.UserRepository {
	return &SqliteRepo{
		db: db,
	}
}
