package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa interface {
	Desativar()
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("\nO ciente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	cliente1 := Client{
		Nome:  "Luciano",
		Idade: 28,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t ", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
	Desativacao(cliente1)
}
