package main

import "fmt"

func main() {
	//i := 0
	//for i < 10 {
	//	time.Sleep(time.Millisecond)

	//	i++
	//	println(i)
	//}

	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementando J: ", j)
	//	time.Sleep(time.Millisecond)
	//}

	//nomes := [3]string{"João", "Davi", "Lucas"}

	//for _, nome := range nomes {
	//	fmt.Println(nome)
	//	time.Sleep(time.Millisecond)
	//}

	//for indice, letra := range "PALAVRA" {
	//	println(indice, string(letra))
	//}

	usuario := map[string]string{
		"nome":      "Jonatas",
		"sobrenome": "Rodrigues",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}


	var i int

	for {
		fmt.Println("Oi")
		i++

		if i == 5 {
			break
		}
	}
}