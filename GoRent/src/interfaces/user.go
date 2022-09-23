package interfaces

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	Save(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, email string) (*models.User, error)
	DeleteUser(email string) (*models.User, error)
}

type UserService interface {
	GetAll() *helpers.Response
	Add(data *models.User) *helpers.Response
	Update(data *models.User, email string) *helpers.Response
	Delete(email string) *helpers.Response
}
