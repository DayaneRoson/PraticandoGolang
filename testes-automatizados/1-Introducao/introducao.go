package main

import (
	"fmt"
	"introducao-testes/endereco"
)

func main() {

	address := endereco.TipoEndereco("Estrada do Nunca")
	fmt.Println(address)

}
