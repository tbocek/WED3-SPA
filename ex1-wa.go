package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "rust/pkg/rust_wasm_bg.wasm")
}

func main() {
	http.HandleFunc("/wasm", handler)
	http.Handle("/", http.FileServer(http.Dir("rust/pkg")))
	http.ListenAndServe(":8080", nil)
}
