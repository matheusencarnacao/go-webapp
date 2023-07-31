package main

import (
	"alura_loja/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
