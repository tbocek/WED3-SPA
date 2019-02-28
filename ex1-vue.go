package main

import (
	//https://godoc.org/github.com/julienschmidt/httprouter, https://github.com/julienschmidt/httprouter
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func vue(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "vue.html")
}

func api_vue(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=10")
	io.Copy(w, res.Body)
}

func main() {
	log.Print("Starting server...")
	router := httprouter.New()
	router.GET("/", vue)
	router.GET("/spa", api_vue)
	log.Fatal(http.ListenAndServe(":8080", router))
}
