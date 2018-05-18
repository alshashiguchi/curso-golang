package main

import (
	"fmt"
)

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo, uma vez que usou as variaveis de retorno o go devolve os valores que estao nela
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
