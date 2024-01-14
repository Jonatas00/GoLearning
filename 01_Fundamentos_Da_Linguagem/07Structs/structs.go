package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero int8
}

func main() {
	fmt.Println("Structs")

	var pessoa usuario
	pessoa.nome = "Jonatas"
	pessoa.idade = 20
	fmt.Println(pessoa)

	enderecoPessoa2 := endereco{"Rua Marília", 21}
	pessoa2 := usuario{"Jonatas", 20, enderecoPessoa2}
	fmt.Println(pessoa2)

	pessoa3 := usuario{idade: 20}
	fmt.Println(pessoa3)
}
