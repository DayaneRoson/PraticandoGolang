package formas

import (
	"math"
)

type Retangulo struct {
	Base   float64
	Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Base
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pow(c.Raio, 2) * math.Pi
}

type Forma interface {
	area() float64
}
