package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CriarToken retorna um token assinado com as permissoes do usuario
func CriarToken(usuarioID string) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true                          // quem tem esse campo esta autorizado ao usar as rotas da api
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // hora atual mais 6 horas em Unix numero 157/848844...
	permissoes["usuarioId"] = usuarioID
	// Token será criado apartir dessas permissões
	// Criando assinatura do token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes) // Criando o token, metodo de assinatura e o HS256
	return token.SignedString([]byte(config.SecretKey))            // Assinatura do token  (secret e a chave para fazer a assinatura)
}

//Autenticar verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // transfroma as permissoes em Map
		return nil
	}
	return errors.New("Token inavalido")
}

func ExtrairusarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%v", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("Token invalido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization") // Extrai o tokeno do Authorization Bearer Token

	// token = Bearer ighbgbgbbg4g5eg51df1g

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return "" // faz a validar quebar
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	// verifica se o metodo usado para gerar o token é valido
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	//Caso seja retornar a chave de verificação
	return config.SecretKey, nil
}
