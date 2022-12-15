package endereco

import "testing"

func TestTipoEndereco(t *testing.T) {
	testAddress := "Avenida Paulista"
	addressReceived := TipoEndereco(testAddress)
	addressExpected := "Avenida"

	if addressExpected != addressReceived {
		t.Errorf("Test on endereco package failed, Expected: %s, received: %s", addressExpected, addressReceived)
	}
}
