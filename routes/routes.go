package routes

import (
	"net/http"

	"github.com/dev-davimarques/go-aplicacao-web-alura/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
