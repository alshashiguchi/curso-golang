package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// evitar funções que causam efeitos colaterais com interface
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// como utilizamos o ponteiro é necessário a utilziação do 	&
	var carro2 esportivo = &ferrari{"F50", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
