package services

import (
	"pc/core/models"
	"pc/core/interfaces"
)

type UserService struct{
	Repo interfaces.UserRepository
}

func (u UserService) NewUser(data models.User){
	if(data.Name == "" || data.Password == ""){
		panic("Error for create user: credentials is missing")
	}
	u.Repo.Save(data)
}
func (u UserService) DeleteUserById(id int) (error){
	if id == 0 {
		panic("Id is null")
	}
	err := u.Repo.Remove(id)
	if err != nil{
		return err
	}
	return nil
}
func (u UserService) EditUser(id int,NewData models.User) error{
	err := u.Repo.Put(id,NewData)
	if err != nil{
		return err
	}
	return nil
}
func (u UserService) FindUser(id int)(models.User,error){
	user,err:=u.Repo.FindOne(id)

	if err != nil{
		return user,err
	}
	return user,nil
}
func (u UserService) ListUsers()([]models.User,error){
	users,err := u.Repo.Find()
	return users,err
}