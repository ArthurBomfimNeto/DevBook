package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	modelos "api/src/models"
	"api/src/repositorios"
	respostas "api/src/resposta"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PostPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairusarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var publicacao modelos.Publicacao

	erro = json.Unmarshal(corpoRequisicao, &publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	publicacao.AutorID = usuarioID

	if erro := publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	repositorio := repositorios.NovoRepositorioDePublicacao(db)
	publicacao.ID, erro = repositorio.CriarPublicacao(publicacao)
	if erro != nil {
		fmt.Println("ENTROU")
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)

}
func GetPublicacoes(w http.ResponseWriter, r *http.Request) {

}
func GetPublicacaoId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	var publicacao modelos.Publicacao

	repositorio := repositorios.NovoRepositorioDePublicacao(db)
	publicacao, erro = repositorio.BuscarPublicacaoID(publicacaoId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)

}
func PutPublicacao(w http.ResponseWriter, r *http.Request) {

}
func DeletePublicacao(w http.ResponseWriter, r *http.Request) {

}
