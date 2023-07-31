package models

import "alura_loja/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := db.ConectaBD()

	selectProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriarProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()

	insertDb, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDb.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeleteProduto(id string) {
	db := db.ConectaBD()

	deleteQuery, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(id)

	defer db.Close()
}

func BuscaProduto(id string) Produto {
	db := db.ConectaBD()

	query, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}

	defer db.Close()

	return p
}

func AtualizaProduto(id, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()

	updateDb, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateDb.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()

}
