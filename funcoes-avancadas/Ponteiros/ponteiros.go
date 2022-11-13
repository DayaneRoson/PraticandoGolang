package main

import (
	"fmt"
)

func inverteSinal(numero int8) int8 {
	return numero * -1
}

func inverteSinalPonteiro(number *int) {
	*number = *number * -1
}

func main() {
	numero := 20
	numeroInvertido := inverteSinal(int8(numero))
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverteSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
