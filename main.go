package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Gary's blog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}


func main() {

	mux := http.NewServeMux()
	mux.Handle("/home", http.HandlerFunc(home))
	mux.Handle("/", http.HandlerFunc(c))

	http.ListenAndServe(":3000", mux)
}