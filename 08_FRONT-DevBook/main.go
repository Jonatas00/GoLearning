package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/internal/router"
	"webapp/internal/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
