package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 20)
	}("Passando parâmetro")

	fmt.Println(retorno)
}