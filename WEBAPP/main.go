package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func init() {
	config.Carregar()
	cookies.ConfigurarCookie()
	utils.CarregarTemplates()
}

func main() {
	r := router.Gerar()

	fmt.Println(fmt.Sprintf("Escutando na porta %v\n", config.Porta))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Porta), r))
}
