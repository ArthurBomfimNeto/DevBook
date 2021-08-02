package controllers

import (
	"net/http"
)

// PostUsuarios insere usuarios no banco
func PostUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuarios"))
}

//GetUsuarios busca por todos usuarios no banco
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuarios"))
}

//GetUsuario busca por um usuario no banco
func GetUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario"))
}

//PutUsuario atualiza um usuario no banco
func PutUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuario"))
}

//DeleteUsuario deleta um usuario no banco
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuario"))
}
