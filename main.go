package main

import (
	"net/http"

	"github.com/dev-davimarques/go-aplicacao-web-alura/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
