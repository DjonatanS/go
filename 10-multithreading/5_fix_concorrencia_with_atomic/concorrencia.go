package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// Variavel a ser acessada por varias threads em paralelo
var guest_counter uint64

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		atomic.AddUint64(&guest_counter, 1)

		w.Write([]byte(fmt.Sprintf("Visitante número %d", guest_counter)))

	})

	http.ListenAndServe(":5020", nil)
}
