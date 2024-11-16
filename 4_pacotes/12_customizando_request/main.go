package main

import (
	"io"
	"net/http"
)

func main() {
	webClient := http.Client{}

	//Criação da request
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	//Http client executando a request
	resp, err := webClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
