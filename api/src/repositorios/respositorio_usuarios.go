package repositorios

import (
	modelos "api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	stmt, erro := repositorio.db.Prepare("insert into usuarios(nome, nick, email, senha) values (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer stmt.Close()

	result, erro := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDinserido, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDinserido), nil
}
