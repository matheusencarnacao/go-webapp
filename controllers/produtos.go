package controllers

import (
	"alura_loja/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco: ", err)
		}

		qtdInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.CriarProduto(nome, descricao, precoFloat, qtdInt)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeleteProduto(idProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	produto := models.BuscaProduto(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco: ", err)
		}

		qtdInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.AtualizaProduto(id, nome, descricao, precoFloat, qtdInt)
	}

	http.Redirect(w, r, "/", 301)
}
