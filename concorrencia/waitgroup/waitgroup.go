package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutines")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutines 3")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutines 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
