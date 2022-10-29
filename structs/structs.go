package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Dayane"
	u.idade = 26
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}
	usuario2 := usuario{"Lucas", 26, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
