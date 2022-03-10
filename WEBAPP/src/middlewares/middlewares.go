package middlewarer

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica a exitencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Caso der erro de não conter o cookie ele redireciona para o login
		_, erro := cookies.Ler(r)
		//fmt.Println(valores, erro)
		if erro != nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		proximaFuncao(w, r)
	}
}
