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

	query := "insert into usuarios(nome, nick, email, senha) values (?,?,?,?)"
	t.Run("sucess", func(t *testing.T) {
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Nome, mockUsuario.Nick,
			mockUsuario.Email, mockUsuario.Senha).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		id, err := repositorio.CriarUser(mockUsuario)

		assert.Nil(t, err)
		assert.NotNil(t, id)

	})

	t.Run("erro-LastInsertID", func(t *testing.T) {
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Nome, mockUsuario.Nick,
			mockUsuario.Email, mockUsuario.Senha).WillReturnResult(sqlmock.NewResult(0, 0))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		id, _ := repositorio.CriarUser(mockUsuario)

		assert.Equal(t, id, uint64(0))
	})

	t.Run("erro-Exec", func(t *testing.T) {
		prep := mock.ExpectPrepare(query)
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

func TestAtualizarUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	t.Run("sucsess", func(t *testing.T) {
		query := `update usuarios set nome = ?, nick = ?, email = ? where id = ?`
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Nome, mockUsuario.Nick,
			mockUsuario.Email, mockUsuario.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.AtualizarUser(mockUsuario, mockUsuario.ID)

		assert.NoError(t, erro)
	})

	t.Run("error-Exec", func(t *testing.T) {
		query := `update usuarios set nome = ?, nick = ?, email = ? where id = ?`
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Nome).WillReturnResult(sqlmock.NewResult(1, 1))

		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.AtualizarUser(mockUsuario, mockUsuario.ID)

		assert.Error(t, erro)
	})

	t.Run("error-Prepare", func(t *testing.T) {
		query := ""
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Nome, mockUsuario.Nick,
			mockUsuario.Email, mockUsuario.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.AtualizarUser(mockUsuario, mockUsuario.ID)

		assert.Error(t, erro)
	})

}

func TestDeletarUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	t.Run("sucsess", func(t *testing.T) {
		query := "delete from usuarios where id = ?"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.ID).WillReturnResult(sqlmock.NewResult(0, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.DeletarUser(mockUsuario.ID)

		assert.NoError(t, erro)
	})

	t.Run("error-Exec", func(t *testing.T) {
		query := "delete from usuarios where id = ?"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs("").WillReturnResult(sqlmock.NewResult(0, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.DeletarUser(mockUsuario.ID)

		assert.Error(t, erro)
	})

	t.Run("error-Prepare", func(t *testing.T) {
		query := ""
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.ID).WillReturnResult(sqlmock.NewResult(0, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.DeletarUser(mockUsuario.ID)

		assert.Error(t, erro)
	})
}

func TestBuscarPorEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "senha"}).
		AddRow(
			mockUsuario.ID,
			mockUsuario.Senha,
		)
	t.Run("success", func(t *testing.T) {
		query := "select id, senha from usuarios where email = ?"
		mock.ExpectQuery(query).WithArgs(mockUsuario.Email).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuario, erro := repositorio.BuscarPorEmail(mockUsuario.Email)

		assert.NotEmpty(t, usuario)
		assert.Nil(t, erro)
	})

	t.Run("erro-Scan", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "senha"}).
			AddRow(
				"erro_forjado",
				mockUsuario.Senha,
			)
		query := "select id, senha from usuarios where email = ?"
		mock.ExpectQuery(query).WithArgs(mockUsuario.Email).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuario, erro := repositorio.BuscarPorEmail(mockUsuario.Email)

		assert.Empty(t, usuario)
		assert.NotNil(t, erro)
	})

	t.Run("erro-query", func(t *testing.T) {
		query := ""
		mock.ExpectQuery(query).WithArgs(mockUsuario.Email).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuario, erro := repositorio.BuscarPorEmail(mockUsuario.Email)

		assert.Empty(t, usuario)
		assert.NotNil(t, erro)
	})

}

func TestSeguir(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	usuario_id := uint64(1)
	seguidor_id := uint64(2)

	t.Run("success", func(t *testing.T) {

		query := "insert ignore into seguidores(usuario_id, seguidor_id) value(?, ?)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(usuario_id, seguidor_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.Seguir(usuario_id, seguidor_id)

		assert.NoError(t, erro)
	})

	t.Run("erro-Exec", func(t *testing.T) {

		query := "insert ignore into seguidores(usuario_id, seguidor_id) value(?, ?)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(usuario_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.Seguir(usuario_id, seguidor_id)

		assert.Error(t, erro)
	})

	t.Run("erro-Prepare", func(t *testing.T) {

		query := ""
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(usuario_id, seguidor_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.Seguir(usuario_id, seguidor_id)

		assert.Error(t, erro)
	})
}

func TestParaDeSeguir(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	usuario_id := uint64(1)
	seguidor_id := uint64(2)

	t.Run("success", func(t *testing.T) {
		query := `delete from seguidores where usuario_id = ? and seguidor_id = ?`
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(usuario_id, seguidor_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.ParaDeSeguir(usuario_id, seguidor_id)

		assert.NoError(t, erro)
	})

	t.Run("error-Exec", func(t *testing.T) {
		query := `delete from seguidores where usuario_id = ? and seguidor_id = ?`
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(usuario_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.ParaDeSeguir(usuario_id, seguidor_id)

		assert.Error(t, erro)
	})

	t.Run("error-Prepare", func(t *testing.T) {
		query := ""
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(usuario_id, seguidor_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.ParaDeSeguir(usuario_id, seguidor_id)

		assert.Error(t, erro)
	})
}

func TestBuscarSeguidores(t *testing.T) {
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

	t.Run("success", func(t *testing.T) {
		usuario_id := uint64(1)

		query := `select u.id, u.nome, u.nick,  u.email, u.criadoEm
		          from usuarios as u 
		          inner join seguidores as s on u.id = s.seguidor_id 
				  where s.usuario_id = ?`

		mock.ExpectQuery(query).WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, erro := repositorio.BuscarSeguidores(usuario_id)

		assert.NotEmpty(t, usuarios)
		assert.NoError(t, erro)
	})

	t.Run("error-Scan", func(t *testing.T) {
		usuario_id := uint64(1)

		rows := sqlmock.NewRows([]string{"id", "nome", "nick", "email", "criadoEm"}).
			AddRow(
				"error-forjado",
				mockUsuario.Nome,
				mockUsuario.Nick,
				mockUsuario.Email,
				mockUsuario.CriadoEm,
			)

		query := `select u.id, u.nome, u.nick,  u.email, u.criadoEm
		          from usuarios as u 
		          inner join seguidores as s on u.id = s.seguidor_id 
				  where s.usuario_id = ?`

		mock.ExpectQuery(query).WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, erro := repositorio.BuscarSeguidores(usuario_id)

		assert.Empty(t, usuarios)
		assert.Error(t, erro)
	})

	t.Run("error-Query", func(t *testing.T) {
		usuario_id := uint64(1)

		query := ""

		mock.ExpectQuery(query).WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, erro := repositorio.BuscarSeguidores(usuario_id)

		assert.Empty(t, usuarios)
		assert.Error(t, erro)
	})

}

func TestBuscarQuemSegue(t *testing.T) {
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
	t.Run("succes", func(t *testing.T) {
		usuario_id := uint64(1)

		query := "select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.usuario_id where seguidor_id =?"

		prep := mock.ExpectPrepare(query)
		prep.ExpectQuery().WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, erro := repositorio.BuscaQuemSegue(usuario_id)

		assert.NotEmpty(t, usuarios)
		assert.Nil(t, erro)
	})

	t.Run("error-Scan", func(t *testing.T) {
		usuario_id := uint64(1)

		rows := sqlmock.NewRows([]string{"id", "nome", "nick", "email", "criadoEm"}).
			AddRow(
				"erro-forjado",
				mockUsuario.Nome,
				mockUsuario.Nick,
				mockUsuario.Email,
				mockUsuario.CriadoEm,
			)

		query := "select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.usuario_id where seguidor_id =?"

		prep := mock.ExpectPrepare(query)
		prep.ExpectQuery().WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, erro := repositorio.BuscaQuemSegue(usuario_id)

		assert.Empty(t, usuarios)
		assert.NotNil(t, erro)
	})

	t.Run("error-query", func(t *testing.T) {
		usuario_id := uint64(1)

		query := "select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.usuario_id where seguidor_id =?"

		prep := mock.ExpectPrepare(query)
		prep.ExpectQuery().WithArgs("").WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, erro := repositorio.BuscaQuemSegue(usuario_id)

		assert.Empty(t, usuarios)
		assert.NotNil(t, erro)
	})

	t.Run("error-Prepare", func(t *testing.T) {
		usuario_id := uint64(1)

		query := ""

		prep := mock.ExpectPrepare(query)
		prep.ExpectQuery().WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		usuarios, erro := repositorio.BuscaQuemSegue(usuario_id)

		assert.Empty(t, usuarios)
		assert.NotNil(t, erro)
	})
}

func TestBuscarSenha(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"senha"}).
		AddRow(
			mockUsuario.Senha,
		)

	t.Run("success", func(t *testing.T) {
		usuario_id := uint64(1)

		query := "select senha from usuarios where id = ?"

		mock.ExpectQuery(query).WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		senha, erro := repositorio.BuscarSenha(usuario_id)

		assert.NotEmpty(t, senha)
		assert.NoError(t, erro)
		assert.Equal(t, senha, mockUsuario.Senha)
	})

	t.Run("error-Query", func(t *testing.T) {
		usuario_id := uint64(1)

		query := ""

		mock.ExpectQuery(query).WithArgs(usuario_id).WillReturnRows(rows)
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		senha, erro := repositorio.BuscarSenha(usuario_id)

		assert.Error(t, erro)
		assert.Empty(t, senha)
	})
}

func TestAtualizarSenha(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	usuario_id := uint64(1)

	t.Run("sucess", func(t *testing.T) {

		query := "update usuarios set senha = ? where id = ?"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Senha, usuario_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.AtualizarSenha(usuario_id, mockUsuario.Senha)

		assert.NoError(t, erro)
	})

	t.Run("error-Exec", func(t *testing.T) {

		query := "update usuarios set senha = ? where id = ?"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Senha).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.AtualizarSenha(usuario_id, mockUsuario.Senha)

		assert.Error(t, erro)
	})

	t.Run("error-Prepare", func(t *testing.T) {

		query := ""
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(mockUsuario.Senha, usuario_id).WillReturnResult(sqlmock.NewResult(1, 1))
		repositorio := repositorios.NovoRepositorioDeUsuarios(db)
		erro := repositorio.AtualizarSenha(usuario_id, mockUsuario.Senha)

		assert.Error(t, erro)
	})
}
