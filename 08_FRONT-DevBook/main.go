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

	url := fmt.Sprintf("%s:%d", config.APPURL, config.Porta)
	fmt.Println(url)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
