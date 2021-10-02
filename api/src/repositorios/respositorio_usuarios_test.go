package repositorios_test

import (
	modelos "api/src/models"
	"api/src/repositorios/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCriarUser(t *testing.T) {
	t.Run("criar user retorno 0", func(t *testing.T) {
		mock := mocks.UsuarioRepositoryMock{}
		usuario, erro := mock.CriarUser(modelos.Usuario{})
		assert.Nil(t, erro)
		assert.Equal(t, uint64(0), usuario)
	})

	t.Run("TESTE DE ERRO", func(t *testing.T) {
		mock := mocks.UsuarioRepositoryMock{}
		_, erro := mock.CriarUser(modelos.Usuario{})
		assert.Error(t, erro)
	})
}

func TestBuscaUsuarioId(t *testing.T) {

}
