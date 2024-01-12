package main

func main() {
	var variavel string = "variavel 1"
	variavel2 := "variavel 2"
	println(variavel)
	println(variavel2)

	var (
		variavel3 string = "l"
		variavel4 string = "b"
	)

	println(variavel3, variavel4)

	variavel5, variavel6 := "312312", "21321321"

	println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	println(constante1)

	variavel5, variavel6 = variavel6, variavel5

	println(variavel5, variavel6)
}