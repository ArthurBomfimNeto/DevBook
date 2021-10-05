package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	modelos "api/src/models"
	"api/src/repositorios"
	usecase "api/src/usecases"
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
func CriarUser(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	usuario.ID, erro = usecase.CriarUser(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	respostas.JSON(w, http.StatusCreated, usuario)
}

//GetUsuarios busca por todos usuarios no banco
func BuscarUser(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusServiceUnavailable, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	usuarios, erro := usecase.BuscarUser(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

//GetUsuario busca por um usuario no banco
func BuscarUserId(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	usuario, erro := usecase.BuscarUserId(id)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

//PutUsuario atualiza um usuario no banco
func AtualizarUser(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	erro = usecase.AtualizarUser(usuario, id)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//DeleteUsuario deleta um usuario no banco
func DeletarUser(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	erro = usecase.DeletarUser(id)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func Seguir(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	erro = usecase.Seguir(usuarioID, seguidorID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func ParaDeSeguir(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	erro = usecase.ParaDeSeguir(usuarioID, seguidorID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	seguidores, erro := usecase.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

func BuscaQuemSegue(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	usuarios, erro := usecase.BuscaQuemSegue(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarSenha(w http.ResponseWriter, r *http.Request) {
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
	usecase := usecase.NovoUsuariosUseCase(repositorio)
	erro = usecase.BuscarSenha(usuarioId, senha)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
