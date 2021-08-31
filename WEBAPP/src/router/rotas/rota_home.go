package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotasHome = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
