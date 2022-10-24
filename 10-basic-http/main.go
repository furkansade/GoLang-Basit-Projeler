package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Anasayfa!")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hakkımda!"))
}

func main() {

	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/hakkimda", AboutPage)
	log.Print("Şu an 9097 portunu dinliyorum!")
	http.ListenAndServe(":9097", nil)
}
