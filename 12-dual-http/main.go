package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/gizli", Incognito)
	http.ListenAndServe(":9099", r)
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, `
	<html><body>
	<h1>Ana Sayfa </h1>
	<a href="/gizli">Gizli Giriş!</a>
	</body></html>
	`)
}

func Incognito(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Kimseye Söyleme!")
}
