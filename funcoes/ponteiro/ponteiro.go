package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // Por valor
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n) // Por referência
	fmt.Println(n)
}
