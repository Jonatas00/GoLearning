package main

import "errors"

func main() {
	var numero int = 1000000000000000000
	println(numero)

	var numero2 uint64 = 050
	println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 12456
	println(numero3)

	//Byte = UINT8
	var numero4 byte = 123
	println(numero4)

	var numeroFloat float32 = 123.45
	println(numeroFloat)

	var numeroFloat2 float64 = 12312313134.5512
	println(numeroFloat2)

	numeroFloat3 := 3.4
	println(numeroFloat3)

	// STRINGS

	var str string = "Texto"
	println(str)

	str2 := "Texto2"
	println(str2)

	char := 'B'
	println(char)

	var stringVazia string
	println(stringVazia)

	var erro error = errors.New("Erro interno")
	println(erro)
}