package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1) // Metodo retorna dois valores o json em slice de bytes e o erro
	fmt.Println(string(p1Json))   // string converte o slice de bytes

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.9,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
	fmt.Println(p2)
}
