package main

import (
	"api/src/config"
	"api/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := routes.Gerar()
	print("git test")
	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
