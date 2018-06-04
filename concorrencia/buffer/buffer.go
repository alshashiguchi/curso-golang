package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou!") // Só executa depois que liberou uma posição no buffer
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch) // Libera uma posição
}
