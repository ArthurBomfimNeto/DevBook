package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func init() {
	utils.CarregarTemplates()
}

func main() {

	r := router.Gerar()

	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
