package interfaces

import modelos "api/src/models"

type PublicacaoRepository interface {
	CriarPublicacao(publicacao modelos.Publicacao) (uint64, error)
	BuscarPublicacaoID(publicacaoId uint64) (modelos.Publicacao, error)
	BuscarPublicacoes(usuarioID uint64) ([]modelos.Publicacao, error)
	AtualizarPublicacao(publicacao modelos.Publicacao, publicacaoID uint64) error
	DeletarPublicacao(publicacaoID uint64) error
	BuscarPublicacoesUsuario(usuarioID uint64) ([]modelos.Publicacao, error)
	Curtir(publicacaoID uint64) error
}
