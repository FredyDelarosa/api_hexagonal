package domain

import "rest/src/ide/domain/entities"

type IDERepository interface {
	GetAll() ([]entities.IDE, error)
	GetByID(id int) (*entities.IDE, error)
	Create(ide *entities.IDE) error
	Update(ide *entities.IDE) error
	Delete(id int) error
}
