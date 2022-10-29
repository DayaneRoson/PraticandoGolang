package main

import "fmt"

func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {

	resultadoSoma, resultadoSubtracao := calculos(30, 50)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
