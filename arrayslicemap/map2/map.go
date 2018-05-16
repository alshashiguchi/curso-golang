package main

import (
	"fmt"
)

func main() {
	// É obrigatorio a virgula no ultimo elemento do map ou a virgula junto com o ultimo elemento ex: "Pedro Junior":   1200.0}
	funcsESalarios := map[string]float64{
		"João José":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
