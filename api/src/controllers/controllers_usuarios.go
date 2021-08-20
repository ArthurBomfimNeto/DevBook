package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	modelos "api/src/models"
	"api/src/repositorios"
	"api/src/seguranca"
	"errors"
	"strconv"
	"strings"

	respostas "api/src/resposta"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// PostUsuarios insere usuarios no banco
func PostUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return

	}

	var usuario modelos.Usuario

	erro = json.Unmarshal(corpoRequisicao, &usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	erro = usuario.Preparar("cadastro")
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.CriarUser(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	respostas.JSON(w, http.StatusCreated, usuario)
}

//GetUsuarios busca por todos usuarios no banco
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusServiceUnavailable, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.BuscarUser(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

//GetUsuario busca por um usuario no banco
func GetUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarUserId(id)

	respostas.JSON(w, http.StatusOK, usuario)

}

//PutUsuario atualiza um usuario no banco
func PutUsuarios(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIdToken, erro := autenticacao.ExtrairusarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
	}

	if usuarioIdToken != id {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel atualizar um usuario sem ser o seu"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	erro = json.Unmarshal(corpoRequisicao, &usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	erro = usuario.Preparar("edicao")
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.AtualizarUser(usuario, id)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//DeleteUsuario deleta um usuario no banco
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIdToken, erro := autenticacao.ExtrairusarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioIdToken != id {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel deletar um usuario sem ser o seu"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.DeletarUser(id)

	respostas.JSON(w, http.StatusNoContent, nil)
}

func PostSeguidor(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	seguidorID, erro := autenticacao.ExtrairusarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.Seguir(usuarioID, seguidorID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func DeleteSeguidor(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	seguidorID, erro := autenticacao.ExtrairusarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel deixar de seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.ParaDeSeguir(usuarioID, seguidorID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func GetSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

func GetQuemSegue(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.BuscaQuemSegue(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func PutSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIdToken, erro := autenticacao.ExtrairusarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioId != usuarioIdToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivél atualizar uma senha que não seja a sua!"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var senha modelos.Senha

	erro = json.Unmarshal(corpoRequisicao, &senha)
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	senhaSalvanoBanco, erro := repositorio.BuscarSenha(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = seguranca.VerificarSenha(senha.Atual, senhaSalvanoBanco)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("A senha atual não condiz com as que esta salva no banco"))
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	erro = repositorio.AtualizarSenha(usuarioId, string(senhaComHash))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
