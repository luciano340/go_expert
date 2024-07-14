package main

import (
	"io"
	"net/http"
)

func main() {
	r, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	res, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	println(string(res))

}
