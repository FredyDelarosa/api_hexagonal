package infrastructure

import (
	"database/sql"
	"fmt"
	"rest/src/auth/domain"
	"rest/src/auth/domain/entities"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) domain.UserRepository {
	return &MySQLUserRepository{db: db}
}

func (repo *MySQLUserRepository) Create(user *entities.User) error {
	query := "INSERT INTO users (email, password, first_name, last_name) VALUES (?, ?, ?, ?)"
	result, err := repo.db.Exec(query, user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return fmt.Errorf("error al crear usuario: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener ID del usuario: %v", err)
	}

	user.ID = int(id)
	return nil
}

func (repo *MySQLUserRepository) GetByEmail(email string) (*entities.User, error) {
	query := "SELECT id, email, password, first_name, last_name, created_at, updated_at FROM users WHERE email = ? LIMIT 1"
	row := repo.db.QueryRow(query, email)

	var user entities.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario: %v", err)
	}

	return &user, nil
}

func (repo *MySQLUserRepository) GetByID(id int) (*entities.User, error) {
	query := "SELECT id, email, password, first_name, last_name, created_at, updated_at FROM users WHERE id = ? LIMIT 1"
	row := repo.db.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario: %v", err)
	}

	return &user, nil
}

func (repo *MySQLUserRepository) Update(user *entities.User) error {
	query := "UPDATE users SET email = ?, password = ?, first_name = ?, last_name = ? WHERE id = ?"
	_, err := repo.db.Exec(query, user.Email, user.Password, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %v", err)
	}
	return nil
}

func (repo *MySQLUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar eliminación: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (repo *MySQLUserRepository) GetAll() ([]entities.User, error) {
	query := "SELECT id, email, password, first_name, last_name, created_at, updated_at FROM users"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %v", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error en la iteración de usuarios: %v", err)
	}

	return users, nil
}
