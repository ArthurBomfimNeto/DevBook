package usecase

import (
	modelos "api/src/models"
	"api/src/repositorios/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCriarUser(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	mockUsuario := modelos.Usuario{
		Nome:  "arthur",
		Nick:  "test",
		Email: "arthur@gmail.com",
		Senha: "123",
	}

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("CriarUser", mockUsuario).Return(uint64(1), nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		id, erro := u.CriarUser(mockUsuario)

		assert.NotEmpty(t, id)
		assert.NoError(t, erro)
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("CriarUser", mockUsuario).Return(uint64(0), errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		id, erro := u.CriarUser(mockUsuario)

		assert.Zero(t, id)
		assert.Error(t, erro)
	})

}

func TestBuscarUser(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	mockResult := &modelos.Usuario{
		ID:       1,
		Nome:     "arthur",
		Nick:     "test",
		Email:    "arthur@gmail.com",
		CriadoEm: time.Time{},
	}
	mockList := []modelos.Usuario{}
	mockList = append(mockList, *mockResult)
	nomeOuNick := "usuario_1"

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarUser", nomeOuNick).Return(mockList, nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuarios, erro := u.BuscarUser(nomeOuNick)

		assert.NotEmpty(t, usuarios) // retorna algo não vazio
		assert.NoError(t, erro)
		assert.Len(t, usuarios, len(mockList)) //verifica pelo tamanho

	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarUser", nomeOuNick).Return(nil, errors.New("Erro")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuarios, erro := u.BuscarUser(nomeOuNick)

		assert.Error(t, erro)
		assert.Empty(t, usuarios)  // retorna algo vazio
		assert.Len(t, usuarios, 0) // verifica pelo tamanho

	})
}

func TestBuscarUserId(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	mockUsuario := modelos.Usuario{
		ID:       1,
		Nome:     "arthur",
		Nick:     "test",
		Email:    "arthur@gmail.com",
		CriadoEm: time.Time{},
	}

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarUserId", mockUsuario.ID).Return(mockUsuario, nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuario, erro := u.BuscarUserId(mockUsuario.ID)

		assert.NotEmpty(t, usuario)
		assert.Nil(t, erro)
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarUserId", mockUsuario.ID).Return(modelos.Usuario{}, errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuario, erro := u.BuscarUserId(mockUsuario.ID)

		assert.Empty(t, usuario)
		assert.Error(t, erro)

	})
}
