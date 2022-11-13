package main

import "fmt"

func recuperarExec() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExec()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 5))
	fmt.Println("Calma, calma não criemos pânico!")

	//PANIC NÃO É TRATAMENTO DE ERRO
}
