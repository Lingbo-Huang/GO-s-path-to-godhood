package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/index", index)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hhhhhh")
	content, _ := os.ReadFile("./index.html")
	w.Write(content)
}
