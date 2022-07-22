package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {

	rotas := rotas.Configurar(mux.NewRouter())

	return rotas
}
