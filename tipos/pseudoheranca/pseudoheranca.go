package main

import (
	"fmt"
)

type carro struct {
	nome             string
	velociadadeAtual int
}

type ferrari struct {
	carro       // campos anonimos - não é herança, na realidade é uma composição
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velociadadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
