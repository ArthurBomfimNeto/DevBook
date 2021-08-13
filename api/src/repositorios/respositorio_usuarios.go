package repositorios

import (
	modelos "api/src/models"
	"database/sql"
	"fmt"
	"reflect"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//Criar insere um usuario no banco de dados
func (repositorio usuarios) CriarUser(usuario modelos.Usuario) (uint64, error) {
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

//Buscar traz todos os usuarios que atendem a um filtro
func (repositorio usuarios) BuscarUser(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	fmt.Println(reflect.TypeOf(usuarios))

	for linhas.Next() {
		var usuario modelos.Usuario
		erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		)
		if erro != nil {

			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}
	return usuarios, nil
}

//Buscar_Id traz um usuario do banco de dados
func (repositorio usuarios) BuscarUserId(id uint64) (modelos.Usuario, error) {

	linha, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		erro = linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		)
		if erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, erro
}

func (repositorio usuarios) AtualizarUser(usuario modelos.Usuario, id uint64) error {
	stmt, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer stmt.Close()

	_, erro = stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, id)
	if erro != nil {
		return erro
	}

	return nil
}

//Deletar exclui as informações de um usuario no banco
func (repositorio usuarios) DeletarUser(id uint64) error {
	stmt, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}

	defer stmt.Close()

	_, erro = stmt.Exec(id)
	if erro != nil {
		return erro
	}

	return nil

}

func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	result, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	var usuario modelos.Usuario

	if result.Next() {
		erro = result.Scan(&usuario.ID, &usuario.Senha)
		if erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	stmt, erro := repositorio.db.Prepare(
		"insert ignore into seguidores(usuario_id, seguidor_id) value(?, ?)")
	if erro != nil {
		return erro
	}

	defer stmt.Close()

	_, erro = stmt.Exec(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) ParaDeSeguir(usuarioID, seguidorID uint64) error {
	stmt, erro := repositorio.db.Prepare(
		`delete from seguidores where usuario_id = ? and seguidor_id = ?`)
	if erro != nil {
		return erro
	}
	_, erro = stmt.Exec(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	rows, erro := repositorio.db.Query(
		`select u.id, u.nome, u.nick,  u.email, u.criadoEm
		 from usuarios as u inner join seguidores as s on u.id = s.seguidor_id where s.usuario_id = ?`, usuarioID)
	if erro != nil {
		return []modelos.Usuario{}, erro
	}

	var usuarios []modelos.Usuario

	var usuario modelos.Usuario

	for rows.Next() {
		erro = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		)
		if erro != nil {
			return []modelos.Usuario{}, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
