package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Dayane", "Roson", 26, 167}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia Elétrica", "UCP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
}
