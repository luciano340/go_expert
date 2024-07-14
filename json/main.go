package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonPuro := []byte(`{"n": 2, "s": 500}`)
	var contax Conta
	err = json.Unmarshal(jsonPuro, &contax)

	if err != nil {
		panic(err)
	}
	println(contax.Saldo)
}
