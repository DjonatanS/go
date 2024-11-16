package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		// Server Side logging
		log.Println("Request processada com sucesso")
		// Client Side logging
		w.Write([]byte("Request processada client side"))
	case <-ctx.Done():
		//Server Side logging
		log.Println("Request cancelada pelo cliente")
		//Client Side logging
		http.Error(w, "Request cancelada pelo client", http.StatusRequestTimeout)
	}
}
