package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handlerSlug(wr http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)["slug"]
	if slug == "" {
		log.Println("slug not provided")
		return
	}
	log.Println("Got slug:", slug)
}

func handlerGetHello(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "Hello Bro\n")
	log.Println("[!] request via", req.Method)
	log.Println(req.URL)
	log.Println(req.Header)
	log.Println(req.Body)
}

func handlerPostEcho(wr http.ResponseWriter, req *http.Request) {
	log.Println("[!] request via:", req.Method)
	log.Println(req.URL)
	log.Println(req.Header)

	/*
		we are going to read it into a buffer as the request body is an io.ReadCloser
		and so we should only readitonce.
	*/
	body, err := ioutil.ReadAll(req.Body)
	log.Println("read >", string(body), "<")
	n, err := io.Copy(wr, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error echoing response", err)
	}
	log.Println("[+] wrote back", n, "bytes")
}

func main() {
	// set some flags for easy debugging
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	// get port from ENV var or default to 9003
	port := "9003"
	if value, exists := os.LookupEnv("SERVER_PORT"); exists {
		port = value
	}

	// off the bat, we can enforce StrictSlash, this is a nice helper function thant means
	// when true, if the route is "/foo/", accessing "/foo" will perform a 301 redirect to the former and vice versa
	// so the app will always see the path as specified in the route.
	// when false: if route path is "/foo", accessing "/foo/" will not match and vice versa

	router := mux.NewRouter().StrictSlash(true)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	router.HandleFunc("/", handlerGetHello).Methods(http.MethodGet)
	router.HandleFunc("/", handlerPostEcho).Methods(http.MethodPost)
	router.HandleFunc("/{slug}", handlerSlug).Methods(http.MethodGet)

	log.Println("[!] starting on", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("[??] couldn't listen and servce", err)
	}
}
