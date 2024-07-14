package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("Teste.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Olá mundo!")
	//se não tiver certeza do que vai colocar no arquivo usar apenas write
	tamanho2, err := f.Write([]byte("\nOlá mundo2!"))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso tamanho: %d bytes\nTamanho2: %d bytes", tamanho, tamanho2)
	f.Close()

	//Leitura
	arquivo, err := os.ReadFile("Teste.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//Leitura em partes - economia de memoria
	arquivo2, err := os.Open("Teste.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
