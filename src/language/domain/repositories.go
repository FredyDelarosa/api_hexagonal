package domain

import "rest/src/language/domain/entities"

type LanguageRepository interface {
	GetAll() ([]entities.Language, error)
	GetByID(id int) (*entities.Language, error)
	Create(language *entities.Language) error
	Update(language *entities.Language) error
	Delete(id int) error
}
