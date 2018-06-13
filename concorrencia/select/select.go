package main

import (
	"fmt"
	"time"

	"github.com/alshashiguchi/titulohtml"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := titulohtml.Titulo(url1)
	c2 := titulohtml.Titulo(url2)
	c3 := titulohtml.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(500 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
	)

	fmt.Println(campeao)
}
