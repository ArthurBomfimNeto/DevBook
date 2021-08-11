package autenticacao

import (
	"api/src/config"
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
