package infrastructureIDE

import (
	"database/sql"
	"rest/src/IDE/domain"
	"rest/src/IDE/domain/entities"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) domain.IDERepository {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Create(ide *entities.IDE) error {
	query := "INSERT INTO ides (name, developer, release_year) VALUES (?, ?, ?)"
	_, err := repo.db.Exec(query, ide.Name, ide.Developer, ide.ReleaseYear)
	return err
}

func (repo *MySQLRepository) GetAll() ([]entities.IDE, error) {
	var ides []entities.IDE
	rows, err := repo.db.Query("SELECT id, name, developer, release_year FROM ides")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ide entities.IDE
		if err := rows.Scan(&ide.ID, &ide.Name, &ide.Developer, &ide.ReleaseYear); err != nil {
			return nil, err
		}
		ides = append(ides, ide)
	}

	return ides, nil
}

func (repo *MySQLRepository) GetByID(id int) (*entities.IDE, error) {
	var ide entities.IDE
	err := repo.db.QueryRow("SELECT id, name, developer, release_year FROM ides WHERE id = ?", id).
		Scan(&ide.ID, &ide.Name, &ide.Developer, &ide.ReleaseYear)

	if err != nil {
		return nil, err
	}

	return &ide, nil
}

func (repo *MySQLRepository) Update(ide *entities.IDE) error {
	query := "UPDATE ides SET name = ?, developer = ?, release_year = ? WHERE id = ?"
	_, err := repo.db.Exec(query, ide.Name, ide.Developer, ide.ReleaseYear, ide.ID)
	return err
}

func (repo *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM ides WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}
