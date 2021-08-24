package repositorios

import (
	"api/src/interfaces"
	modelos "api/src/models"
	"database/sql"
	"fmt"
)

type publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacao(db *sql.DB) interfaces.PublicacaoRepository {
	return &publicacoes{db}
}

func (repositorio publicacoes) CriarPublicacao(publicacao modelos.Publicacao) (uint64, error) {
	stmt, erro := repositorio.db.Prepare(
		"insert into publicacoes(titulo, conteudo, autor_id) values (?, ?, ?)")
	if erro != nil {
		fmt.Println("ENTROU1")

		return 0, erro
	}

	result, erro := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if erro != nil {
		fmt.Println("ENTROU2")

		return 0, erro
	}

	usuarioId, erro := result.LastInsertId()
	if erro != nil {
		fmt.Println("ENTRO3")

		return 0, erro
	}
	return uint64(usuarioId), nil
}

func (repositorio publicacoes) BuscarPublicacaoID(publicacaoId uint64) (modelos.Publicacao, error) {
	stmt, erro := repositorio.db.Prepare(
		`select p.*, u.nick  
		 from publicacoes p 
		 inner join usuarios u on p.autor_id = u.id where p.id = ?`)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	row, erro := stmt.Query(publicacaoId)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer row.Close()

	var publicacao modelos.Publicacao
	if row.Next() {
		erro = row.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		)
		if erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

func (repositorio publicacoes) BuscarPublicacoes(usuarioID uint64) ([]modelos.Publicacao, error) {
	stmt, erro := repositorio.db.Prepare(` 
	select distinct p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id 
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc;`)
	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

	rows, erro := stmt.Query(usuarioID, usuarioID)
	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

	var publicacoes []modelos.Publicacao

	for rows.Next() {
		var publicacao modelos.Publicacao
		erro = rows.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		)
		if erro != nil {
			return []modelos.Publicacao{}, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio publicacoes) AtualizarPublicacao(publicacao modelos.Publicacao, publicacaoID uint64) error {
	stmt, erro := repositorio.db.Prepare(`update publicacoes 
	                                           set titulo = ?,
											       conteudo = ? 
												   where id = ?`)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio publicacoes) DeletarPublicacao(publicacaoID uint64) error {
	stmt, erro := repositorio.db.Prepare(`delete from publicacoes where id = ?`)
	if erro != nil {
		return erro
	}

	defer stmt.Close()

	_, erro = stmt.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio publicacoes) BuscarPublicacoesUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	stmt, erro := repositorio.db.Prepare(
		`select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where u.id = ?`)
	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

	defer stmt.Close()

	rows, erro := stmt.Query(usuarioID)
	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

	defer rows.Close()

	var publicacoes []modelos.Publicacao

	for rows.Next() {
		var publicacao modelos.Publicacao
		erro = rows.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		)
		if erro != nil {
			return []modelos.Publicacao{}, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio publicacoes) Curtir(publicacaoID uint64) error {
	stmt, erro := repositorio.db.Prepare(
		`update publicacoes set curtidas = curtidas + 1 where id = ?`)
	if erro != nil {
		return erro
	}

	defer stmt.Close()

	_, erro = stmt.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}
