package main

import "fmt"

func generica(interf any) {
	fmt.Println(interf)
}

func main() {
	generica(35353)

	mapa := map[any] any {
		1: "String",
		float32(100): true,
		"String": "String",
	}
	fmt.Println(mapa)
}
