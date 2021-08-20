package interfaces

import modelos "api/src/models"

type PublicacaoRepository interface {
	CriarPublicacao(publicacao modelos.Publicacao) (uint64, error)
	BuscarPublicacaoID(publicacaoId uint64) (modelos.Publicacao, error)
}
