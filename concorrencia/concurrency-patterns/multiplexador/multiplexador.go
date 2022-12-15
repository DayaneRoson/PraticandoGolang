package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo"), escrever("Golang"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(input1, input2 <-chan string) <-chan string {
	output := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-input1:
				output <- mensagem
			case mensagem := <-input2:
				output <- mensagem
			}
		}
	}()
	return output
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
