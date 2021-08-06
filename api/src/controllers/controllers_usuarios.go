package controllers

import (
	//"api/src/banco"
	"api/src/banco"
	modelos "api/src/models"
	"api/src/repositorios"
	"fmt"
	"strings"

	//"api/src/repositorios"
	respostas "api/src/resposta"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// PostUsuarios insere usuarios no banco
func PostUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	fmt.Println(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return

	}

	var usuario modelos.Usuario

	erro = json.Unmarshal(corpoRequest, &usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	erro = usuario.Preparar()
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	erro = json.NewEncoder(w).Encode(usuario)
	respostas.JSON(w, http.StatusCreated, usuario)
}

//GetUsuarios busca por todos usuarios no banco
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusServiceUnavailable, erro)
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
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
