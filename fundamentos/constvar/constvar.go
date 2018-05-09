package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduziada de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunverência é", area)

	// Bloco de declaração de variáveis
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// Declaração de variáveis e atribuição de valore em uma única linha
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
