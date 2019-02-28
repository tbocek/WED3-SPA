package main

import (
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "market-ws.html")
}

func proxy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=10")
	io.Copy(w, res.Body)
}

func ws(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	messageType, _, _ := conn.NextReader()
	for {
		res, _ := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=10")
		w, _ := conn.NextWriter(messageType)
		io.Copy(w, res.Body)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	log.Print("Starting server...")
	router := httprouter.New()
	router.GET("/", test)
	router.GET("/api", proxy)
	router.GET("/wss", ws)
	log.Fatal(http.ListenAndServe(":8080", router))
}
