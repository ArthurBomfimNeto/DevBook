package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")
	config.Carregar()
	r := router.Gerar()

	fmt.Print(config.StringConexao)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
