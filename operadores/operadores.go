package main

import "fmt"

func main() {

	//ARITMETICOS
	adicao := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(adicao, subtracao, divisao, multiplicacao, restoDaDivisao)

	var n1 int8 = 10
	var n2 int16 = 20

	soma := n1 + int8(n2)
	fmt.Println(soma)

	//ATRIBUICAO
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	fmt.Println("-----------------")
	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("-----------------")
	//OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
}
