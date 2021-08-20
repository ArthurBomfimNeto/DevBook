package mocks

import (
	modelos "api/src/models"
	"errors"
)

type RepositoryMocks struct {
	LancaErro bool
}

func (r RepositoryMocks) CriarUser(usuario modelos.Usuario) (uint64, error) {
	if r.LancaErro {
		return 0, errors.New("ERRO!")
	}

	return 0, nil
}

func (r RepositoryMocks) BuscarUser(nomeOuNick string) ([]modelos.Usuario, error) {
	return []modelos.Usuario{}, nil
}

func (r RepositoryMocks) BuscarUserId(id uint64) (modelos.Usuario, error) {
	return modelos.Usuario{}, nil
}

func (r RepositoryMocks) AtualizarUser(usuario modelos.Usuario, id uint64) error {
	return nil
}
func (r RepositoryMocks) DeletarUser(id uint64) error {
	return nil
}

func (r RepositoryMocks) BuscarPorEmail(email string) (modelos.Usuario, error) {
	return modelos.Usuario{}, nil
}

func (r RepositoryMocks) Seguir(usuarioID, seguidorID uint64) error {
	return nil
}

func (r RepositoryMocks) ParaDeSeguir(usuarioID, seguidorID uint64) error {
	return nil
}

func (r RepositoryMocks) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	return []modelos.Usuario{}, nil
}
