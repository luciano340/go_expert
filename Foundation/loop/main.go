package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	num := []int{1, 2, 3, 4, 5}

	for k, v := range num {
		println(k, v)
	}

	for {
		println("Infinito")
	}
}
