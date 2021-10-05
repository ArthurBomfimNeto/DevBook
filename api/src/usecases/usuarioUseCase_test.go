package usecase

import (
	modelos "api/src/models"
	"api/src/repositorios/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestAtualizaruser(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	mockUsuario := modelos.Usuario{
		Nome:  "arthur",
		Nick:  "test",
		Email: "arthur@gmail.com",
		Senha: "123",
	}

	id := uint64(1)

	t.Run("sucsess", func(t *testing.T) {
		mockUsuarioRepository.On("AtualizarUser", mockUsuario, id).Return(nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.AtualizarUser(mockUsuario, id)

		assert.Nil(t, erro)
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("AtualizarUser", mockUsuario, id).Return(errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.AtualizarUser(mockUsuario, id)

		assert.Error(t, erro)
	})

}

func TestDeletaruser(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	id := uint64(1)

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("DeletarUser", id).Return(nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.DeletarUser(id)

		assert.Nil(t, erro)
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("DeletarUser", id).Return(errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.DeletarUser(id)

		assert.Error(t, erro)
	})
}

func TestSeguir(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	usuario_id := uint64(1)
	seguidor_id := uint64(2)

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("Seguir", usuario_id, seguidor_id).Return(nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.Seguir(usuario_id, seguidor_id)

		assert.NoError(t, erro)
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("Seguir", usuario_id, seguidor_id).Return(errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.Seguir(usuario_id, seguidor_id)

		assert.Error(t, erro)
	})

}

func TestParaDeSeguir(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	usuario_id := uint64(1)
	seguidor_id := uint64(2)

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("ParaDeSeguir", usuario_id, seguidor_id).Return(nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.ParaDeSeguir(usuario_id, seguidor_id)

		assert.NoError(t, erro)
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("ParaDeSeguir", usuario_id, seguidor_id).Return(errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.ParaDeSeguir(usuario_id, seguidor_id)

		assert.Error(t, erro)
	})
}

func TestBuscarSeguidores(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	mockUsuario := modelos.Usuario{
		Nome:  "arthur",
		Nick:  "test",
		Email: "arthur@gmail.com",
		Senha: "123",
	}

	mockList := []modelos.Usuario{}
	mockList = append(mockList, mockUsuario)

	usuario_id := uint64(1)

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarSeguidores", usuario_id).Return(mockList, nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuarios, erro := u.BuscarSeguidores(usuario_id)

		assert.NotEmpty(t, usuarios)
		assert.NoError(t, erro)
		assert.Len(t, usuarios, len(mockList))
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarSeguidores", usuario_id).Return([]modelos.Usuario{}, errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuarios, erro := u.BuscarSeguidores(usuario_id)

		assert.Empty(t, usuarios)
		assert.Error(t, erro)
		assert.Len(t, usuarios, 0)
	})

}

func TestBuscaQuemSegue(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	mockUsuario := modelos.Usuario{
		Nome:  "arthur",
		Nick:  "test",
		Email: "arthur@gmail.com",
		Senha: "123",
	}

	mockList := []modelos.Usuario{}
	mockList = append(mockList, mockUsuario)

	usuario_id := uint64(1)

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("BuscaQuemSegue", usuario_id).Return(mockList, nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuarios, erro := u.BuscaQuemSegue(usuario_id)

		assert.NotEmpty(t, usuarios)
		assert.NoError(t, erro)
		assert.Len(t, usuarios, len(mockList))
	})

	t.Run("error", func(t *testing.T) {
		mockUsuarioRepository.On("BuscaQuemSegue", usuario_id).Return([]modelos.Usuario{}, errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		usuarios, erro := u.BuscaQuemSegue(usuario_id)

		assert.Empty(t, usuarios)
		assert.Error(t, erro)
		assert.Len(t, usuarios, 0)
	})

}

func TestBuscarSenha(t *testing.T) {
	mockUsuarioRepository := mocks.UsuarioRepository{}
	senha := modelos.Senha{
		Atual: "123456",
		Nova:  "1234",
	}
	usuario := modelos.Usuario{
		Senha: "$2a$10$fUQjonJnfSLP.iV7LBVHKO1B425ar/REuHkFnvnCeCXRxXU/mIKfa",
	}

	usuario_id := uint64(1)

	t.Run("success", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarSenha", usuario_id).Return(usuario.Senha, nil).Once()
		mockUsuarioRepository.On("AtualizarSenha", usuario_id, mock.AnythingOfType("string")).Return(nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.BuscarSenha(usuario_id, senha)

		assert.NoError(t, erro)
	})

	t.Run("error-AtualizarSenhaRepository", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarSenha", usuario_id).Return(usuario.Senha, nil).Once()
		mockUsuarioRepository.On("AtualizarSenha", usuario_id, mock.AnythingOfType("string")).Return(errors.New("")).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.BuscarSenha(usuario_id, senha)

		assert.Error(t, erro)
	})

	t.Run("error-VerificarSenha", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarSenha", usuario_id).Return("usuario.Senha", nil).Once()
		mockUsuarioRepository.On("AtualizarSenha", usuario_id, mock.AnythingOfType("string")).Return(nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.BuscarSenha(usuario_id, senha)

		assert.Error(t, erro)
	})

	t.Run("error-BuscarSenhaRepository", func(t *testing.T) {
		mockUsuarioRepository.On("BuscarSenha", usuario_id).Return("", errors.New("")).Once()
		mockUsuarioRepository.On("AtualizarSenha", usuario_id, mock.AnythingOfType("string")).Return(nil).Once()
		u := NovoUsuariosUseCase(&mockUsuarioRepository)
		erro := u.BuscarSenha(usuario_id, senha)

		assert.Error(t, erro)
	})

}
