package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	myServer := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))
	myServer.ListenAndServe()

}
