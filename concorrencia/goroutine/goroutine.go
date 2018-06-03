package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Executando de forma serial
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1)

	// Executando como goroutine
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}
