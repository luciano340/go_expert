package main

import "fmt"

type Conta struct {
	saldo float64
}

func (c Conta) simular(valor float64) float64 {
	c.saldo += valor
	return c.saldo
}

func (c *Conta) adicionar(valor float64) float64 {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{
		saldo: 200,
	}
	conta.simular(200)
	fmt.Printf("O saldo da conta é %v\n", conta.saldo)

	conta.adicionar(150)
	fmt.Printf("O saldo da conta é %v\n", conta.saldo)
}
