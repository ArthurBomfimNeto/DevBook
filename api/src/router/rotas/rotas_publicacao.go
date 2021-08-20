package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacao = []Rota{
	{
		URI:                "/publicacao",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PostPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetPublicacaoId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.PutPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletePublicacao,
		RequerAutenticacao: true,
	},
}
