package main

import (
	"fmt"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				//time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	// Termina o laço for, se o channel ficar aberto o laço entra em um loop infinito ou gerar deadlock
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c) // cap - retorna a capacidade do objeto

	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
