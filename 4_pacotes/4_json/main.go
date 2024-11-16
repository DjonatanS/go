package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero string `json:"numero"`
	Saldo  int    `json:"saldo"`
}

func main() {
	conta := Conta{Numero: "5050", Saldo: 2350}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	//Json sempre vai retornar em bytes
	print(string(res))

	//O encoder vai jogar output para o stdout
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)
	if err != nil {
		println(err)
	}

	jsonData := []byte(`{"numero": "50", "saldo": 500}`)
	var new_conta Conta
	err = json.Unmarshal(jsonData, &new_conta)
	if err != nil {
		println(err)
	}
	println(new_conta.Saldo)
}
