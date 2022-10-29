package main

import "fmt"

func main() {

	fmt.Println("Ponteiros")
	var variavel1 int8 = 10
	var variavel2 int8 = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	fmt.Println("--------------")

	var variavel3 int8 = 100
	var ponteiro *int8

	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	fmt.Println("--------------")

	variavel3 = 110
	fmt.Println(variavel3, *ponteiro) //desreferenciacao

}
