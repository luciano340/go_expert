package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(2, 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(valor)

}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A SOMA É MAIOR QUE 50")
	}
	return a + b, nil

}
