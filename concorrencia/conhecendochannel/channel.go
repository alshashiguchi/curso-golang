package main

import (
	"fmt"
)

func main() {
	// channel é um tipo em go assim como o int
	ch := make(chan int, 1) //Criando um canal de comunicação será enviado e recebido inteiros

	ch <- 1 // Enviando dados para o canal (escrita)
	<-ch    // Recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
