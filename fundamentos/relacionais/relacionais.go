package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("3 != 2:", 3 != 2)
	fmt.Println("3 < 2:", 3 < 2)
	fmt.Println("3 > 2:", 3 > 2)
	fmt.Println("3 <= 2:", 3 <= 2)
	fmt.Println("3 >= 2:", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1, d2)
	fmt.Println("Datas (=):", d1 == d2)      // Comparando os valores e não a referencia da memoria
	fmt.Println("Datas equal", d1.Equal(d2)) // Comparando os valores e não a referencia da memoria

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas:", p1 == p2) // Comparando os valores e não a referencia da memoria
}
