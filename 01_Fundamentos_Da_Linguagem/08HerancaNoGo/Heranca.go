package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa1 := pessoa{"João", "Pedro", 20, 179}

	fmt.Println(pessoa1)

	e1 := estudante{pessoa1, "Engenharia", "USP"}

	fmt.Println(e1)

	e2 := estudante{pessoa{"André", "Carlos", 23, 168}, "Enfermagem", "PUC"}
	fmt.Println(e2)
}