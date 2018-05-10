package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // Não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando....")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// laco infinito
	for {
		fmt.Println("Para Sempre...")
		time.Sleep(time.Second * 5) // imprime a cada 5 segundo
	}
}
