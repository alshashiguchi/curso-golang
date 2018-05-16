package main

import (
	"fmt"
)

func main() {

	// var aprovados map[int] string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Arthur"
	aprovados[987654321] = "Pedro"
	aprovados[122535478] = "Maria"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[122535478])
	delete(aprovados, 122535478)
	fmt.Println(aprovados[122535478])
}
