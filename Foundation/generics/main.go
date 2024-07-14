package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Luciano": 250, "Taynara": 850}
	m2 := map[string]float64{"Luciano": 2.50, "Taynara": 8.50}
	m3 := map[string]MyNumber{"Luciano": 250, "Taynara": 850}
	println(soma(m))
	println(soma(m2))
	println(soma(m3))
	println(compara(10, 10))
}
