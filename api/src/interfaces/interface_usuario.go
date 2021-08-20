package interfaces

import modelos "api/src/models"

type UsuarioRepository interface {
	CriarUser(usuario modelos.Usuario) (uint64, error)
	BuscarUser(nomeOuNick string) ([]modelos.Usuario, error)
	BuscarUserId(id uint64) (modelos.Usuario, error)
	AtualizarUser(usuario modelos.Usuario, id uint64) error
	DeletarUser(id uint64) error
	BuscarPorEmail(email string) (modelos.Usuario, error)
	Seguir(usuarioID, seguidorID uint64) error
	ParaDeSeguir(usuarioID, seguidorID uint64) error
	BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error)
	BuscaQuemSegue(usuarioId uint64) ([]modelos.Usuario, error)
	BuscarSenha(usuarioId uint64) (string, error)
	AtualizarSenha(usuarioId uint64, senhaComHash string) error
}
