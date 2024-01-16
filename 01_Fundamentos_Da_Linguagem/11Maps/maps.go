package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Jonatas",
		"sobrenome": "Rodrigues",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string {
		"nome" : {
			"primeiro": "Jonatas",
			"segundo": "Rodrigues",
		},
		"curso": {
			"nome": "ADS",
			"campus": "Uninter",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}