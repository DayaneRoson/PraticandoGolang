package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int = 100000000000000
	number := 333333333333333333
	fmt.Println(numero, number)

	//int8, int16, int32, int64 = inteiros (precisa especificar os bits)
	//uint8, uint16, uint32, uint64 = inteiros sem sinal (negativo não pode, mesma lógica dos bits)

	/*alias
	int32 = rune
	uint8 = byte (obviamente porque o go é uma linguagem de baixo nível)
	*/

	//NUMEROS REAIS
	var apelidoInt32 rune = 123456
	fmt.Println(apelidoInt32)

	var float float32 = 0.123
	fmt.Println(float)
	float64 := 123344556.444
	fmt.Println(float64)

	//Strings
	var str string
	fmt.Println(str)

	var dar string
	fmt.Println(dar)

	string := "Teste"
	fmt.Println(string)

	char := 'D'
	fmt.Println(char)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}
