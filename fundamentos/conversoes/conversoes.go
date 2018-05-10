package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado
	fmt.Println("Teste " + string(97)) // Codigo correspondente na tabela ascii

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para boolean
	b, _ := strconv.ParseBool("true") // "1" = True qualquer outro valor retorna false

	if b {
		fmt.Println("Verdadeiro")
	}
}
