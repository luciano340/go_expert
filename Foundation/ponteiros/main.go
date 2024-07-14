package main

func soma(a, b *int) int {
	*a = 60
	*b = 60
	return *a + *b
}
func main() {
	var1 := 10
	var2 := 120

	soma(&var1, &var2)
	println(var1)
	println(var2)

}
