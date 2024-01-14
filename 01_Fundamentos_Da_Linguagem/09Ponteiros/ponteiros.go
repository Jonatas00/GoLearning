package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++

	fmt.Println(variavel1, variavel2)

	// ponteiro é uma referência da memória
	var variavel3 int = 100
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro) // desreferenciação -> *

	variavel3 = 50

	fmt.Println(variavel3, *ponteiro)
}