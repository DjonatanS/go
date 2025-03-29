package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Variavel a ser acessada por varias threads em paralelo
var guest_counter uint64

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Não permite a variavel ser acessada por mais de uma thread de forma simultânea
		m.Lock()
		guest_counter++
		m.Unlock()

		w.Write([]byte(fmt.Sprintf("Visitante número %d", guest_counter)))

	})

	http.ListenAndServe(":5020", nil)
}
