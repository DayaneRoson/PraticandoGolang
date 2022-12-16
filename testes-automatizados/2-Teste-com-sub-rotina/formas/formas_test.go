package formas_test

import (
	"math"
	. "teste-com-sub-rotina/formas"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Test failed using Retangulo. Expect: %f, found: %f", areaEsperada, areaRecebida)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Test failed using Circulo. Expect: %f, found: %f", areaEsperada, areaRecebida)
		}
	})
}
