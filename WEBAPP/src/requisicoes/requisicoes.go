package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutenticacao é utilizada para colocar o token  na requisição
// Recebe a requisição feita pela webb e cria uma nova chmando a API com o token no Header
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	// apenas cria a requisição porém não envia ainda a API
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}

	//Chapar o token no Authorization - Bearer token como feito manual no postman

	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	//Criar um client para realizar a requisição na API
	client := &http.Client{}
	//faz a requisição com .Do e é retornado uma response
	resposne, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}
	//retorna a response com as publicações extraida do banco
	return resposne, nil
}
