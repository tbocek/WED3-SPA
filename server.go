package main

import (
	"encoding/json"
	"strconv"
	//https://godoc.org/github.com/julienschmidt/httprouter, https://github.com/julienschmidt/httprouter
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type Model struct {
	Name string
}

var counter = 0

func nospa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	template, err := template.ParseFiles("index.gtpl")
	if err != nil {
		log.Fatal(err)
	}
	model := Model{Name: ps.ByName("name") + ", counter=" + strconv.Itoa(counter)}
	err = template.Execute(w, &model)
	if err != nil {
		log.Fatal(err)
	}
	counter++
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "index.html")
}

func spa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(Model{Name: ps.ByName("name")})
}

func vue(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "vue.html")
}

func main() {
	log.Print("Starting server...")
	router := httprouter.New()
	router.GET("/nospa/:name", nospa)
	router.GET("/spa/:name", spa)
	router.GET("/", index)
	router.GET("/vue", vue)
	log.Fatal(http.ListenAndServe(":8080", router))
}
