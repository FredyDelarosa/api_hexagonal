package infrastructure

import (
	"database/sql"
	"rest/src/language/domain"
	"rest/src/language/domain/entities"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) domain.LanguageRepository {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Create(language *entities.Language) error {
	query := "INSERT INTO languages (name, paradigm, created_year) VALUES (?, ?, ?)"
	_, err := repo.db.Exec(query, language.Name, language.Paradigm, language.CreatedYear)
	return err
}

func (repo *MySQLRepository) GetAll() ([]entities.Language, error) {
	var languages []entities.Language
	rows, err := repo.db.Query("SELECT id, name, paradigm, created_year FROM languages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lang entities.Language
		if err := rows.Scan(&lang.ID, &lang.Name, &lang.Paradigm, &lang.CreatedYear); err != nil {
			return nil, err
		}
		languages = append(languages, lang)
	}

	return languages, nil
}

func (repo *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM languages WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *MySQLRepository) Update(language *entities.Language) error {
	query := "UPDATE languages SET name = ?, paradigm = ?, created_year = ? WHERE id = ?"
	_, err := repo.db.Exec(query, language.Name, language.Paradigm, language.CreatedYear, language.ID)
	return err
}

func (repo *MySQLRepository) GetByID(id int) (*entities.Language, error) {
	var language entities.Language
	query := "SELECT id, name, paradigm, created_year FROM languages WHERE id = ?"
	err := repo.db.QueryRow(query, id).Scan(&language.ID, &language.Name, &language.Paradigm, &language.CreatedYear)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &language, nil
}
