package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	m = make(map[string]string)
)

func get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uuid := ps.ByName("uuid")
	w.Header().Set("Content-Type", "application/json")
	if len(m[uuid]) > 0 {
		io.Copy(w, strings.NewReader(m[uuid]))
	} else {
		w.Write([]byte("[]"))
	}
}

func post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uuid := ps.ByName("uuid")

	if b, err := ioutil.ReadAll(r.Body); err == nil {
		m[uuid] = string(b)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	log.Print("Starting server...")
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "portfolio/index.html")
	})
	router.ServeFiles("/dist/*filepath", http.Dir("portfolio/dist"))

	router.GET("/portfolio/:uuid", get)
	router.POST("/portfolio/:uuid", post)

	log.Fatal(http.ListenAndServe(":8080", router))
}
