package application

import (
	"fmt"
	"rest/src/auth/domain"
	"rest/src/auth/domain/entities"
	"rest/src/core/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	repo domain.UserRepository
}

type LoginResponse struct {
	User  *entities.User `json:"user"`
	Token string         `json:"token"`
}

func NewLoginUseCase(repo domain.UserRepository) *LoginUseCase {
	return &LoginUseCase{repo: repo}
}

func (uc *LoginUseCase) Execute(email, password string) (*LoginResponse, error) {
	// Validar que el email no esté vacío
	if email == "" {
		return nil, fmt.Errorf("el email es requerido")
	}

	// Validar que la contraseña no esté vacía
	if password == "" {
		return nil, fmt.Errorf("la contraseña es requerida")
	}

	// Buscar usuario por email
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error al buscar usuario: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("credenciales inválidas")
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("credenciales inválidas")
	}

	// Generar token JWT
	token, err := utils.GenerateToken(user.ID, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return nil, fmt.Errorf("error al generar token: %v", err)
	}

	// No retornar la contraseña
	user.Password = ""

	return &LoginResponse{
		User:  user,
		Token: token,
	}, nil
}
