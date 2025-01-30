package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=postgres password=12345 host=localhost sslmode=diseble"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

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
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
