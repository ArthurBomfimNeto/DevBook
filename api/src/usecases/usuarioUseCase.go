package usecase

import (
	"api/src/interfaces"
	modelos "api/src/models"
)

type usuarioUseCase struct {
	usuarioRepository interfaces.UsuarioRepository
}

func NovoUsuariosUseCase(u interfaces.UsuarioRepository) interfaces.UsuariosUseCase {
	return &usuarioUseCase{
		usuarioRepository: u,
	}
}

func (u usuarioUseCase) CriarUser(usuario modelos.Usuario) (uint64, error) {

	ultimoIDinserido, erro := u.usuarioRepository.CriarUser(usuario)
	if erro != nil {
		return 0, erro
	}
	return ultimoIDinserido, nil
}

func (u *usuarioUseCase) BuscarUser(nomeOuNick string) ([]modelos.Usuario, error) {
	usuarios, erro := u.usuarioRepository.BuscarUser(nomeOuNick)
	if erro != nil {
		return []modelos.Usuario{}, erro
	}

	return usuarios, nil
}

func (u *usuarioUseCase) BuscarUserId(id uint64) (modelos.Usuario, error) {
	usuario, erro := u.usuarioRepository.BuscarUserId(id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	return usuario, nil
}
