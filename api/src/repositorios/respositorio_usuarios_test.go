package repositorios_test

import (
	modelos "api/src/models"
	"api/src/repositorios"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var mockUsuario = modelos.Usuario{
	ID:       1,
	Nome:     "arthur",
	Nick:     "test",
	Email:    "arthur@gmail.com",
	Senha:    "1234",
	CriadoEm: time.Time{},
}

func TestCriarUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("An Error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	queryInsert := "insert into usuarios(nome, nick, email, senha) values (?,?,?,?)"
	t.Run("sucess", func(t *testing.T) {
		prep := mock.ExpectPrepare(queryInsert)
		prep.ExpectExec().WithArgs(mockUsuario.Nome, mockUsuario.Nick,
			mockUsuario.Email, mockUsuario.Senha).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		id, err := repositorio.CriarUser(mockUsuario)

		assert.Nil(t, err)
		assert.NotNil(t, id)

	})

	t.Run("erro-LastInsertID", func(t *testing.T) {
		prep := mock.ExpectPrepare(queryInsert)
		prep.ExpectExec().WithArgs(mockUsuario.Nome, mockUsuario.Nick,
			mockUsuario.Email, mockUsuario.Senha).WillReturnResult(sqlmock.NewResult(0, 0))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		id, _ := repositorio.CriarUser(mockUsuario)

		assert.Equal(t, id, uint64(0))
	})

	t.Run("erro-Exec", func(t *testing.T) {
		prep := mock.ExpectPrepare(queryInsert)
		prep.ExpectExec().WithArgs("").WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		id, err := repositorio.CriarUser(mockUsuario)

		assert.Error(t, err)
		assert.Equal(t, id, uint64(0))
	})

	t.Run("erro-Prepare", func(t *testing.T) {
		prep := mock.ExpectPrepare("")
		prep.ExpectExec().WithArgs(mockUsuario.Nome, mockUsuario.Nick,
			mockUsuario.Email, mockUsuario.Senha).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		id, err := repositorio.CriarUser(mockUsuario)

		assert.Error(t, err)
		assert.Equal(t, id, uint64(0))
	})

}

func TestBuscarUser(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nome", "nick", "email", "criadoEm"}).
		AddRow(
			mockUsuario.ID,
			mockUsuario.Nome,
			mockUsuario.Nick,
			mockUsuario.Email,
			mockUsuario.CriadoEm,
		)

	t.Run("sucess", func(t *testing.T) {

		usuarioName := fmt.Sprintf("%%%s%%", mockUsuario.Nome)

		query := "select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?"
		mock.ExpectQuery(query).WithArgs(usuarioName, usuarioName).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuario, err := repositorio.BuscarUser(mockUsuario.Nome)

		assert.NoError(t, err)
		assert.NotNil(t, usuario)
	})

	t.Run("error-Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "nome", "nick", "email", "criadoEm"}).
			AddRow(
				"teste",
				mockUsuario.Nome,
				mockUsuario.Nick,
				mockUsuario.Email,
				mockUsuario.CriadoEm,
			)
		usuarioName := fmt.Sprintf("%%%s%%", mockUsuario.Nome)
		query := "select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?"
		mock.ExpectQuery(query).WithArgs(usuarioName, usuarioName).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, err := repositorio.BuscarUser(mockUsuario.Nome)

		assert.Error(t, err)
		assert.Equal(t, usuarios, []modelos.Usuario{})

	})

	t.Run("error-query", func(t *testing.T) {

		usuarioName := fmt.Sprintf("%%%s%%", mockUsuario.Nome)

		query := ""
		mock.ExpectQuery(query).WithArgs(usuarioName, usuarioName).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, err := repositorio.BuscarUser(mockUsuario.Nome)

		assert.Error(t, err)
		assert.Equal(t, usuarios, []modelos.Usuario{})
	})

}

func TestBuscarUserID(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nome", "nick", "email", "criadoEm"}).
		AddRow(
			mockUsuario.ID,
			mockUsuario.Nome,
			mockUsuario.Nick,
			mockUsuario.Email,
			mockUsuario.CriadoEm,
		)

	t.Run("sucess", func(t *testing.T) {

		query := "select id, nome, nick, email, criadoEm from usuarios where id = ?"
		mock.ExpectQuery(query).WithArgs(mockUsuario.ID).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuario, err := repositorio.BuscarUserId(mockUsuario.ID)

		assert.NoError(t, err)
		assert.NotNil(t, usuario)
	})

	t.Run("erro-Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "nome", "nick", "email", "criadoEm"}).
			AddRow(
				"teste",
				mockUsuario.Nome,
				mockUsuario.Nick,
				mockUsuario.Email,
				mockUsuario.CriadoEm,
			)

		query := "select id, nome, nick, email, criadoEm from usuarios where id = ?"
		mock.ExpectQuery(query).WithArgs(mockUsuario.ID).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuario, err := repositorio.BuscarUserId(mockUsuario.ID)

		assert.Error(t, err)
		assert.Equal(t, usuario, modelos.Usuario{})
	})

	t.Run("erro-query", func(t *testing.T) {
		query := ""
		mock.ExpectQuery(query).WithArgs(mockUsuario.ID).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuario, err := repositorio.BuscarUserId(mockUsuario.ID)

		assert.Error(t, err)
		assert.Equal(t, usuario, modelos.Usuario{})
	})

}
