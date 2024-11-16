package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursoTemplate := template.Must(template.New("content.html").ParseFiles(templates...))
		err := cursoTemplate.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 60},
			{"Python", 45},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)

}
