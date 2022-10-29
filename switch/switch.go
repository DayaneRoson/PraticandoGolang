package main

import "fmt"

func diaDaSemana(numero int8) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "sábado"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int8) string {
	var dia string

	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough
	case numero == 2:
		dia = "Segunda"
	case numero == 3:
		dia = "Terça"
	case numero == 4:
		dia = "Quarta"
	case numero == 5:
		dia = "Quinta"
	case numero == 6:
		dia = "Sextou"
	case numero == 7:
		dia = "Sábado"
	default:
		dia = "Número inválido"
	}

	return dia
}

func main() {

	fmt.Println("Switch")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
