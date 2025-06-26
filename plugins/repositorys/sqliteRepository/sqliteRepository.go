package sqliterepository

import (
	"database/sql"
	"pc/core/interfaces"
	"pc/core/models"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepo struct {
	db *sql.DB
}

func (s *SqliteRepo) Save(user models.User) error {
	_, err := s.db.Exec(`INSERT INTO users(name, password) VALUES (?, ?)`, user.Name, user.Password)
	return err
}
func  (s *SqliteRepo) Find() ([]models.User,error){
    rows,err := s.db.Query(`SELECT id,name,password FROM users`)
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id,&user.Name,&user.Password)
		if err != nil{
		   return nil,err
		}
		users = append(users, user)
	}
	return users,err
}

func (s *SqliteRepo) Put(id int,NewData models.User) error{
	user,err:= s.FindOne(id)
	if err != nil{
		return err
	}
	NewDataOfUser := NewData

	if NewDataOfUser.Name == ""{
		NewDataOfUser.Name = user.Name
	}
	if NewDataOfUser.Password == ""{
		NewDataOfUser.Password = user.Password
	}
	_,err= s.db.Exec(`UPDATE users SET name=?,password=? WHERE id=?`,NewDataOfUser.Name,NewDataOfUser.Password,user.Id)
	return err
}


func (s  *SqliteRepo) Remove(id int) error{
	_,err:=s.db.Exec(`DELETE FROM users WHERE id=?`,id)
	if err != nil{
		return err
	}
	return nil
}
func (s *SqliteRepo) FindOne(id int) (models.User, error) {
	var user models.User
	row := s.db.QueryRow(`SELECT id, name, password FROM users WHERE id = ?`, id)
	err := row.Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SqliteRepoInitialize(db *sql.DB) interfaces.UserRepository {
	return &SqliteRepo{
		db: db,
	}
}
