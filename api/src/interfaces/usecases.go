package interfaces

import modelos "api/src/models"

type UsuariosUseCase interface {
	BuscarUser(nomeOuNick string) ([]modelos.Usuario, error)
	CriarUser(usuario modelos.Usuario) (uint64, error)
	BuscarUserId(id uint64) (modelos.Usuario, error)
}
