package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Println("9098 portu dinleniyor!")
	http.ListenAndServe(":9098", nil)
}
