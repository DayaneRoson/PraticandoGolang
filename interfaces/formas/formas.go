package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	base   float64
	altura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.base
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %.2f\n", f.area())
}

func main() {

	r := retangulo{10, 15}
	escreverArea(r)
	c := circulo{43}
	escreverArea(c)
}
