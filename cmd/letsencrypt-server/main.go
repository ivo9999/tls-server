package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "it's working")
}

func main() {
	http.HandleFunc("/", index)
	err := http.Serve(autocert.NewListener("server-name"), nil)
	if err != nil {
		log.Fatal("error: ", err)
	}
}
