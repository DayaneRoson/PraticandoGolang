package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type gato struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	g := gato{"Felicia", "Gato Norueguês das Neves", 2}
	//fmt.Println(g)
	inputJson, error := json.Marshal(g)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(inputJson)
	cat := bytes.NewBuffer(inputJson)
	fmt.Println(cat)

	c2 := map[string]string{
		"nome": "Leninha",
		"raca": "Ragdoll",
	}

	inputMapJson, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(error)
	}

	fmt.Println(bytes.NewBuffer(inputMapJson))
}
