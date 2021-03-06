package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42) // can only pass in 1 peice of data
	if err != nil {
		log.Fatalln(err)
	}
}
