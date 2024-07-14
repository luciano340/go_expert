package main

import "fmt"

const a = "ftm_tipagem"

type ID int

var (
	b bool
	c int
	d string
	e float64
	f ID = 10
)

func main() {
	var my_array [3]int
	my_array[0] = 100
	my_array[1] = 200
	my_array[2] = 300

	fmt.Println(my_array[len(my_array)-1])

	for i, v := range my_array {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}

}
