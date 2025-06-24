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
func (u UserService) DeleteUserById(){
}
func (u UserService) EditUser(){

}
func (u UserService) FindUser(){

}
func (u UserService) ListUsers()([]models.User,error){
	users,err := u.Repo.Find()
	return users,err
}