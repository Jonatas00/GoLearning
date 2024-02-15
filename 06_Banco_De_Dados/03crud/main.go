package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CREATE READ UPDATE DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuario", servidor.BuscarUSuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUSuario).Methods(http.MethodGet)

	fmt.Println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
