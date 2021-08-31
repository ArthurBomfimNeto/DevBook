package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	modelos "webapp/src/models"
	"webapp/src/respostas"
)

//Fazer login faz a requisição a API para realizar o login do usuario
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Analisa o corpo da requisição e preenche o r.Form com os dados retornados

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	fmt.Println(bytes.NewBuffer(usuario)) // transforma o slice de bytes em json legivel

	url := fmt.Sprintf("%v/login", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	var DadosAutenticacao modelos.DadosAutenticacao
	erro = json.NewDecoder(response.Body).Decode(&DadosAutenticacao) // recebe um json de decodifica ele para um struct e afins
	if erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
	}

	if erro = cookies.Salvar(w, DadosAutenticacao.ID, DadosAutenticacao.Token); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}
