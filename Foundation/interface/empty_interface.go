package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Luciano"
	ShowType(x)
	ShowType(y)
}

func ShowType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor %v\n", t, t)
}
