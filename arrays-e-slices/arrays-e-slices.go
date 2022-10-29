package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Arrays e Slices")
	var array1 [5]int8

	array1[4] = 4
	fmt.Println(array1)

	fmt.Println("-----------")
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4"}

	fmt.Println(array2)

	fmt.Println("-----------")
	array3 := [...]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	//SLICES
	fmt.Println("-----------")
	slice := []int8{48}
	fmt.Println(slice)

	fmt.Println("-----------")
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 27)
	fmt.Println("new slice")
	fmt.Println(slice)

	slice2 := array3[3:4]
	fmt.Println(slice2)

	//ARRAYS INTERNOS
	fmt.Println("-----------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	fmt.Println("-----------")
	slice3 = append(slice3, 4, 6)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 12)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
