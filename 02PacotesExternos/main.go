package main

import (
	"fmt"
	"modulo/pacoteAux"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("123")

	fmt.Println(erro)
}