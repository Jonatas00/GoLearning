package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência != Paralelismo
	go escrever("Olá mundo")
	escrever("Programando em GO")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}