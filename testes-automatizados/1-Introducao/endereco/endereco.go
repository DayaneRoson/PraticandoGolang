package endereco

import "strings"

func TipoEndereco(endereco string) string {
	validType := []string{"Rua", "Servidao", "Avenida", "Estrada", "Praça"}

	firstWord := strings.Split(endereco, " ")[0]

	validAddress := false

	for _, varType := range validType {
		if strings.ToLower(varType) == strings.ToLower(firstWord) {
			validAddress = true
		}

		if validAddress {
			return firstWord
		}
	}
	return "Invalid Type"
}
