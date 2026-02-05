package domain

import "rest/src/auth/domain/entities"

type UserRepository interface {
	GetByEmail(email string) (*entities.User, error)
	GetByID(id int) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id int) error
	GetAll() ([]entities.User, error)
}
