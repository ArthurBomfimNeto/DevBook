package usecase

import (
	"api/src/interfaces"
	modelos "api/src/models"
	"api/src/seguranca"
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

func (u *usuarioUseCase) AtualizarUser(usuario modelos.Usuario, id uint64) error {
	erro := u.usuarioRepository.AtualizarUser(usuario, id)
	if erro != nil {
		return erro
	}

	return nil
}

func (u *usuarioUseCase) DeletarUser(id uint64) error {
	erro := u.usuarioRepository.DeletarUser(id)
	if erro != nil {
		return erro
	}

	return nil
}

func (u *usuarioUseCase) Seguir(usuarioID, seguidorID uint64) error {
	erro := u.usuarioRepository.Seguir(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

func (u *usuarioUseCase) ParaDeSeguir(usuarioID, seguidorID uint64) error {
	erro := u.usuarioRepository.ParaDeSeguir(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

func (u *usuarioUseCase) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	seguidores, erro := u.usuarioRepository.BuscarSeguidores(usuarioID)
	if erro != nil {
		return []modelos.Usuario{}, erro
	}

	return seguidores, nil
}

func (u *usuarioUseCase) BuscaQuemSegue(usuarioId uint64) ([]modelos.Usuario, error) {
	usuarios, erro := u.usuarioRepository.BuscaQuemSegue(usuarioId)
	if erro != nil {
		return []modelos.Usuario{}, erro
	}
	return usuarios, nil
}

func (u *usuarioUseCase) BuscarSenha(usuarioId uint64, senha modelos.Senha) error {
	senhaSalvanoBanco, erro := u.usuarioRepository.BuscarSenha(usuarioId)
	if erro != nil {
		return erro
	}
	erro = seguranca.VerificarSenha(senha.Atual, senhaSalvanoBanco)
	if erro != nil {
		return erro
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		return erro
	}

	erro = u.usuarioRepository.AtualizarSenha(usuarioId, string(senhaComHash))
	if erro != nil {
		return erro
	}
	return nil
}
