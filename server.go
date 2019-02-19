package main

import (
	//https://godoc.org/github.com/julienschmidt/httprouter, https://github.com/julienschmidt/httprouter
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type Model struct {
	Name string
}

func helloWorld(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	template, err := template.ParseFiles("index.gtpl")
	if err != nil {
		log.Fatal(err)
	}
	model := Model{Name: ps.ByName("name")}
	err = template.ExecuteTemplate(w, "index.gtpl", &model)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Print("Starting server...")
	router := httprouter.New()
	router.GET("/nospa/:name", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", router))
}
