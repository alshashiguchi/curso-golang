package main

import (
	"fmt"
)

func main() {
	// ... compilador conta quantidade de elementos no array
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // se não colocar o _ ele vai passar o indice para a variavel num
		fmt.Println(num)
	}
}
