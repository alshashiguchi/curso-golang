package main

import "fmt"

func notaParaConceito(n float64) string {

	value := int(n)
	switch {
	case value >= 9 && value <= 10:
		return "A"
	case value >= 8 && value <= 9:
		return "B"
	case value >= 5 && value <= 8:
		return "C"
	case value >= 3 && value <= 8:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
