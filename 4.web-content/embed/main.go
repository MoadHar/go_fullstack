package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version/version.txt
	version string

	//go:embed static/*
	staticEmbed embed.FS

	//go:embed tmpl/*.html
	tmplEmbed embed.FS
)

type staticHandler struct {
	staticPath string
	indexPage  string
}

func (h staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	log.Println(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)
	log_fileinfo, err := os.Stat(path)
	log.Println("| >> | log_fileinfo:", log_fileinfo)

	log.Print("[!] using embed mode")
	fsys, err := fs.Sub(staticEmbed, "static")
	if err != nil {
		panic(err)
	}
	http.FileServer(http.FS(fsys)).ServeHTTP(w, r)
}

// rederFiles renders file and push data (d) into the template to be rendered
func rendreFiles(tmpl string, w http.ResponseWriter, d interface{}) {
	t, err := template.ParseFS(tmplEmbed, fmt.Sprintf("tmpl/%s.html", tmpl))
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	result := "Login "
	r.ParseForm()

	if validateUser(r.FormValue("username"), r.FormValue("password")) {
		result += "successfull"
	} else {
		result += "unsuccessfull"
	}
	rendreFiles("msg", w, result)
}

func validateUser(username, password string) bool {
	return (username == "admin") && (password == "admin")
}

func main() {
	log.Println("Server version:", version)
	router := mux.NewRouter()
	router.HandleFunc("/login", postHandler).Methods("POST")

	spa := staticHandler{
		staticPath: "static",
		indexPage:  "index.html",
	}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":9639",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
/bin/bash: line 1: :w: command not found
