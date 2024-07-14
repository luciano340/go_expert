package main

import "fmt"

//composição
type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

//Criando tipo
type Client2 struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("\nO ciente %s foi desativado", c.Nome)
}

func main() {

	cliente1 := Client{
		Nome:  "Luciano",
		Idade: 28,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t ", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
	cliente1.Desativar()
}
