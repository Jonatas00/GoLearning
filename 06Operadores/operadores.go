package main

func main() {
	// Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 1 / 2
	multiplicacao := 1 * 2
	restoDaDivisao := 10 % 3

	println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + numero2
	println(soma2)

	// Atribuição
	var variavel string = "String"
	variavel2 := "String2"

	println(variavel, variavel2)

	// Relacionais

	println(1 > 2)
	println(1 >= 2)
	println(1 < 3)
	println(1 <= 2)
	println(1 == 1)
	println(3 != 2)

	// Logicos

	verdadeiro, falso := true, false
	println(verdadeiro && falso) // and
	println(verdadeiro || falso) // or
	println(!verdadeiro) // not

	// Unarios

	numero := 10
	numero++
	println(numero)
	numero--
	println(numero)
	numero += 15
	println(numero)
	numero -= 15
	println(numero)
	numero *= 3
	println(numero)
	numero /= 3
	println(numero)


}