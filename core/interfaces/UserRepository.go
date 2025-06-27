package interfaces

import (
	"pc/core/models"
)

type UserRepository interface {
	Save(user models.User) error
	Remove(id int64) error
	Put(id int, NewData models.User) error
	FindOne(id int) (models.User, error)
	Find() ([]models.User,error)
}

