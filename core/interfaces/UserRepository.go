package interfaces

import (
	"pc/core/models"
)

type UserRepository interface {
	Save(user models.User) error
	Remove(id int) error
	// Put(id int, NewData any) error
	//FindOne(id int) error
	Find() ([]models.User,error)
}

