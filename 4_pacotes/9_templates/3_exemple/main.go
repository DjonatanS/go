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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursoTemplate := template.Must(template.New("template.html").ParseFiles("template.html"))
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
