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
	fmt.Printf("O f é do tipo %T - valor %v", f, f)
	fmt.Printf("O e é do tipo %T - valor %v", e, e)
}
