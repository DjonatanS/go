package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/blog", Blog{Title: "Olá, blog!"}.ServeHTTP)
	http.ListenAndServe(":8080", mux)

}

type Blog struct {
	Title string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func (blog Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(blog.Title))
}
