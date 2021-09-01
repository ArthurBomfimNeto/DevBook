package rotas

import (
	"net/http"
	middlewarer "webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

//Configurar cria todas as rotas e retorna elas
func Configurar(router *mux.Router) *mux.Router {
	rotas := RotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, RotasHome)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI,
				middlewarer.Logger(middlewarer.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI,
				middlewarer.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}

	}

	fileServer := http.FileServer(http.Dir("./assets/")) //
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
