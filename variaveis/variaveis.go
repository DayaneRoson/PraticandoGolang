package main

import "fmt"

func main() {
	var nominho string = "Leninha"
	gatos := "Branquinha"

	fmt.Println(nominho)
	fmt.Println(gatos)

	var (
		cor   string = "Vermelha"
		fruta string = "Lichia"
	)
	fmt.Println(cor, fruta)

	racaGato, apelido := "Gato Noruegues", "Gato das Neves"
	fmt.Println(racaGato, apelido)

	const frase = "Eu sou uma constante imutável muahahaha!"

	fmt.Println(frase)

	racaGato, apelido = apelido, racaGato
	fmt.Println(racaGato, apelido)
}
