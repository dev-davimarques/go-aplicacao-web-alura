package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
	EmEstoque       bool
}

var temp = template.Must(template.ParseGlob("template/*.html"))

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{"Camiseta", "Azul, bem bonita", 39, 5, true},
		{"Tênis", "Confortável", 89, 3, true},
		{"Fone", "Muito bom", 59, 2, true},
		{"Produto novo", "Muito legal", 1.99, 1, true},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
