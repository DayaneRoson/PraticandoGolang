package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá Mundo!")
	}()

	retorno := func(text string) string {
		return fmt.Sprintf("recebido -> %s", text)
	}("Day")

	fmt.Println(retorno)

	func(number int) {
		fmt.Println(number)
	}(54)
}
