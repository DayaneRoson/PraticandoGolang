package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// 	time.Sleep(2 * time.Second)
	// }

	// nomes := [3]string{"Day", "Lucas", "Livea"}

	// for _, value := range nomes {
	// 	fmt.Println(value)
	// }

	// for key, value := range "PALAVRA" {
	// 	fmt.Println(key, string(value))
	// }

	usuario := map[string]string{
		"nome":      "Go",
		"sobrenome": "lang",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for { //loop infinito
		fmt.Println("E aí ??")
	}
}
