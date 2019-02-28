package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "market.html")
}

func api(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=10")
	io.Copy(w, res.Body)
}

func main() {
	log.Print("Starting server...")
	router := httprouter.New()
	router.GET("/", root)
	router.GET("/api", api)
	log.Fatal(http.ListenAndServe(":8080", router))
}
