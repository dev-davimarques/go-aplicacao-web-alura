package controllers

import (
	"html/template"
	"net/http"

	"github.com/dev-davimarques/go-aplicacao-web-alura/models"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
