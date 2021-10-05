package interfaces

import modelos "api/src/models"

type UsuariosUseCase interface {
	BuscarUser(nomeOuNick string) ([]modelos.Usuario, error)
	CriarUser(usuario modelos.Usuario) (uint64, error)
	BuscarUserId(id uint64) (modelos.Usuario, error)
	AtualizarUser(usuario modelos.Usuario, id uint64) error
	DeletarUser(id uint64) error
	Seguir(usuarioID, seguidorID uint64) error
	ParaDeSeguir(usuarioID, seguidorID uint64) error
	BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error)
	BuscaQuemSegue(usuarioId uint64) ([]modelos.Usuario, error)
	BuscarSenha(usuarioId uint64, senha modelos.Senha) error
}
