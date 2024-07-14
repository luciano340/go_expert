package main

import "fmt"

func main() {
	salarios := map[string]int{"Lucaino": 2, "salada": 200}
	fmt.Printf("%d\n", salarios["Lucaino"])
	delete(salarios, "Luciano")

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário s é %d\n", salario)
	}
}
