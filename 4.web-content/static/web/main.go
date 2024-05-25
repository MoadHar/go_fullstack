package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("starting server on port 3456...")
	err := http.ListenAndServe(":3456", nil)
	if err != nil {
		log.Fatal("[-error] starting up server: ", err)
	}
}
