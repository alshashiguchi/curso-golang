package main

import (
	"fmt"

	"github.com/alshashiguchi/titulohtml"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := multiplexar(
		titulohtml.Titulo("https://www.google.com.br", "https://www.uol.com.br"),
		titulohtml.Titulo("https://www.amazon.com.br", "https://www.youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
