package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a URL para comunicação com a API
	APIURL string
	// Porta onde a aplicação web está rodando
	Porta = 0
	// HashKey é utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
)

func Carregar() {
	var erro error

	// Carrega as variaveis de ambient que conte no arquivo .env
	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")

	Porta, erro = strconv.Atoi(os.Getenv("PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
