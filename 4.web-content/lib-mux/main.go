package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handlerGetHello(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "Hello // univer bro\n")
	log.Println(req.Method)
	log.Println(req.URL)
	log.Println(req.Header)
	log.Println(req.Body)
}

type foo struct{}

func (f foo) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "Hello // alternate // bro\n")
	log.Println(req.Method)
	log.Println(req.URL)
	log.Println(req.Header)
	log.Println(req.Body)
}

func main() {
	// set some flags for easy debugging
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	port := "9077"
	// get a port number from ENV or default to 9077
	if value, exists := os.LookupEnv("SERVER_PORT"); exists {
		port = value
	}

	// we could use default mux then use http.handleFunc and http.ListenAndServe(port,nil) but
	// he finds it best to create our own, as we'll see in later patterns we'll cover graceful
	// shutdoiwn as well
	router := http.NewServeMux()

	srv := http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20, // This good to limit how much data to accept
	}

	// This to show an alternate way to declare a handler by having a struct that implements
	// the ServetHTTP(...) interface

	dummyHandler := foo{}

	router.HandleFunc("/", handlerGetHello)
	router.Handle("/1", dummyHandler)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Couldn't ListenAndServ()", err)
	}
}
