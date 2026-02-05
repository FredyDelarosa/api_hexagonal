package application

import (
	"fmt"
	"rest/src/auth/domain"
	"rest/src/auth/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase struct {
	repo domain.UserRepository
}

func NewRegisterUseCase(repo domain.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{repo: repo}
}

func (uc *RegisterUseCase) Execute(email, password, firstName, lastName string) (*entities.User, error) {
	// Validar que el email no esté vacío
	if email == "" {
		return nil, fmt.Errorf("el email es requerido")
	}

	// Validar que la contraseña no esté vacía
	if password == "" {
		return nil, fmt.Errorf("la contraseña es requerida")
	}

	// Validar que la contraseña tenga al menos 6 caracteres
	if len(password) < 6 {
		return nil, fmt.Errorf("la contraseña debe tener al menos 6 caracteres")
	}

	// Verificar que el email no existe
	existingUser, _ := uc.repo.GetByEmail(email)
	if existingUser != nil {
		return nil, fmt.Errorf("el email ya está registrado")
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error al encriptar la contraseña: %v", err)
	}

	// Crear el usuario
	user := &entities.User{
		Email:     email,
		Password:  string(hashedPassword),
		FirstName: firstName,
		LastName:  lastName,
	}

	// Guardar en la base de datos
	if err := uc.repo.Create(user); err != nil {
		return nil, fmt.Errorf("error al registrar el usuario: %v", err)
	}

	// No retornar la contraseña
	user.Password = ""

	return user, nil
}
