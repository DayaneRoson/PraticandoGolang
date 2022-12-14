package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Livea",
		"sobrenome": "Roson",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "das Neves",
		},
		"curso": {
			"nome":   "Letras",
			"campus": "kkk",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["pets"] = map[string]string{
		"nome": "Leninha",
		"raca": "gato",
	}

	fmt.Println(usuario2)
}
