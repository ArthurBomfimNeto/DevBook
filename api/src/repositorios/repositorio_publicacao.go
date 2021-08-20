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
		"select titulo, conteudo, autor_id, curtidas, criadoEm from publicacoes where id = ?")
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
		row.Scan(
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
		)
	}
	return publicacao, nil
}
