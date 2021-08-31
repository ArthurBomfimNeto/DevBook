package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroAPI struct {
	Erro string `json:"erro"`
}

//JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json") // sempre manter json pois a API trafega dados json
	w.WriteHeader(statusCode)

	erro := json.NewEncoder(w).Encode(dados)
	if erro != nil {
		log.Fatal(erro)
	}
}

//TratarStatusCodeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeErro(w http.ResponseWriter, r *http.Response) { //r e o response em json que a api retorna
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro) // transforma um json em struct
	JSON(w, r.StatusCode, erro)
}
