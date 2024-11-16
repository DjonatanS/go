package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{Nome: "Curso de Go", CargaHoraria: 40}
	cursoTemplate := template.Must(template.New("CursoTemplate").Parse("Nome: {{.Nome}}\nCarga Horária: {{.CargaHoraria}} horas\n"))
	err := cursoTemplate.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
