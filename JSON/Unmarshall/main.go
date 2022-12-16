package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type gato struct {
	Nome  string `json:"-"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	inputJson := `{"nome":"Frajola","raca":"Sem Raça definida","idade":4}`
	var g gato
	if erro := json.Unmarshal([]byte(inputJson), &g); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(g)

	input2Json := `{"nome":"Felix","raca":"Siames"}`

	g2 := make(map[string]string)
	if error := json.Unmarshal([]byte(input2Json), &g2); error != nil {
		log.Fatal(error)
	}

	fmt.Println(g2)
}
