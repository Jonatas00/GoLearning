package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1[5] int
	array1[4] = 2
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)
	
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 16)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)
	slice2 = append(slice2, "Posição nova!!")

	fmt.Println(slice2)

	// Arrays internos

	slice3 := make([]int, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 33)
	slice3 = append(slice3, 34)
	slice3 = append(slice3, 34)

	fmt.Println("Tamanho: ", len(slice3))
	fmt.Println("Capacidade: ", cap(slice3))
	
	fmt.Println(slice3)

	slice4 := make([]int, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 32)
	fmt.Println("Tamanho: ", len(slice4))
	fmt.Println("Capacidade:", cap(slice4))
	fmt.Println(slice4)
}
