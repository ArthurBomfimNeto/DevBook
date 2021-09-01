package cookies

import (
	"errors"
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

// Tipo Utilizado para codificar e decofificar os dados colocado no browser
var s *securecookie.SecureCookie

func ConfigurarCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Salvar(w http.ResponseWriter, ID, token string) error {

	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	// Codificar dados para salvar no cookie
	dadosCodificados, erro := s.Encode("dados", dados) // primeiro parametro nome do cookie segundo os dados em si
	if erro != nil {
		return errors.New("ENTROJU AQUI")
	}

	//Responsavel por colocar o cookie no browsier
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Ler(r *http.Request) (map[string]string, error) {
	// pega o map do cookie id e token
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	// descodificação dos campos

	valores := make(map[string]string)

	erro = s.Decode("dados", cookie.Value, &valores)
	if erro != nil {
		return nil, erro
	}
	//map com os valores decodificado
	return valores, nil
}
