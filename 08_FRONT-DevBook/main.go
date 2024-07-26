package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/internal/config"
	"webapp/internal/cookies"
	"webapp/internal/router"
	"webapp/internal/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()

	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d \n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
