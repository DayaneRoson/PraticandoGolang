package endereco

import "testing"

type testingScenes struct {
	addressInput    string
	addressExpected string
}

func TestTipoEndereco(t *testing.T) {

	testingScenes := []testingScenes{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada da Saudade", "Estrada"},
		{"Servidao de Nárnia", "Servidao"},
		{"Praça das Rosas", "Praça"},
		{" ", "Invalid Type"},
	}

	for _, scene := range testingScenes {
		addressReceived := TipoEndereco(scene.addressInput)

		if addressReceived != scene.addressExpected {
			t.Errorf("Test on endereco package failed: Expected: %s, found: %s", scene.addressExpected, addressReceived)
		}
	}
}
