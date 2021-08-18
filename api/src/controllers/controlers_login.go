package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	modelos "api/src/models"
	respostas "api/src/resposta"
	"api/src/seguranca"
	"api/src/usecase/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var usuario modelos.Usuario

	erro = json.Unmarshal(corpoRequisicao, &usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	// USE CASE
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioDobanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = seguranca.VerificarSenha(usuario.Senha, usuarioDobanco.Senha)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// w.Write([]byte("Usuario logado!"))
	// USE CASE
	token, _ := autenticacao.CriarToken(fmt.Sprint(usuarioDobanco.ID))
	w.Write([]byte(token))
}
