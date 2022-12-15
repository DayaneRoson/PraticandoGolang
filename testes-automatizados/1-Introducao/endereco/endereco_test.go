package endereco_test

import (
	. "introducao-testes/endereco"
	"testing"
)

type testingScenes struct {
	addressInput    string
	addressExpected string
}

func TestTipoEndereco(t *testing.T) {
	t.Parallel()

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

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Test failed")
	}
}
