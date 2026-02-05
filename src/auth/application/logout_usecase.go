package application

type LogoutUseCase struct {
	// En una implementación real, aquí iría la lógica para invalidar tokens
	// Por ahora es un logout simple que confirma la desconexión
}

func NewLogoutUseCase() *LogoutUseCase {
	return &LogoutUseCase{}
}

func (uc *LogoutUseCase) Execute(userID int) error {
	// Aquí podrías implementar lógica para:
	// - Invalidar JWT tokens
	// - Eliminar sesiones
	// - Registrar el logout en auditoría
	// Por ahora simplemente confirmamos que fue exitoso
	return nil
}
