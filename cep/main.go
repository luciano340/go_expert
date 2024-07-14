package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}

		var data cep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println(data)

		file, err := os.Create("cidade.txt")

		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprint("CEP: %s", data.Cep))
	}
}
