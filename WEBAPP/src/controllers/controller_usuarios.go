package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//CriarUsuario chama A API para cadastrar um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Analisa o corpo da requisição e atualiza o r.Form

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(usuario)) // transforma o slice de bytes em json legivel

	// realizando a requisição na API com os dados extraido da app web
	response, erro := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		log.Fatal(erro)
	}

	defer response.Body.Close()

	fmt.Println(response.Body)
}
