package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{"Buddha", "The belief of no beliefs"}
	gandhi := sage{"Gandhi", "Be the change"}
	mlk := sage{"Martin Luthar King", "Hatred never wins"}

	sages := []sage{buddha, gandhi, mlk}

	err := tmpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
