package main

import "fmt"

func main() {
	_slice := []int{2, 4, 6, 8, 10}
	fmt.Printf("\n\nlen=%d, cap=%d\n%v", len(_slice), cap(_slice), _slice)
	fmt.Printf("\n\nlen=%d, cap=%d\n%v", len(_slice[:0]), cap(_slice[:0]), _slice[:0])
	fmt.Printf("\n\nlen=%d, cap=%d\n%v", len(_slice[:4]), cap(_slice[:4]), _slice[:4])
	fmt.Printf("\n\nlen=%d, cap=%d\n%v", len(_slice[2:]), cap(_slice[2:]), _slice[2:])

	s := append(_slice, 250)
	fmt.Printf("\n\nlen=%d, cap=%d\n%v", len(s), cap(s), s)
}
