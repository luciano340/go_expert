package main

import "fmt"

func main() {

	total := func() int {
		return sum(1, 555, 666, 4874, 81415) * 2
	}()

	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
