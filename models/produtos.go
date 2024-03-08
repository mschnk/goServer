package models

import (
	_ "github.com/lib/pq"
	"github.com/mschnk/goServer/db"
	_ "github.com/mschnk/goServer/db"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
		db := db.ConectaComBancoDeDados()
		selectDeTodosOsProdutos, err := db.Query("select * from produtos")
		if err != nil {
			panic(err.Error())
		}
		p := Produto{}
		produtos := []Produto{}
	
		for selectDeTodosOsProdutos.Next() {
			//declarando variaveis para receber os valores do banco depois
			var id, quantidade int
			var nome, descricao string
			var preco float64
			//passando os valores do banco para as variaveis, caso nao tenha valores, ele retorna um erro
			err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
			if err != nil {
				panic(err.Error())
			}
			//passando os valores das variaveis para a struct de p: Produto
			p.Nome = nome
			p.Descricao = descricao
			p.Preco = preco
			p.Quantidade = quantidade
			//adicionando a struct p na lista de produtos
			produtos = append(produtos, p)
		}
		defer db.Close()
		return produtos
}