package main

import (
	"net/http"
	"text/template"

	"github.com/dev-davimarques/go-aplicacao-web-alura/models"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
