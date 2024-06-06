package main

import (
	chap5 "chap5/gen"
	"chap5/pkg"
	"context"
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version/version.txt
	version string

	//go:embed static
	staticEmbed embed.FS

	//go:embed css/*
	cssEmbed embed.FS

	//go:embed tmpl/*.html
	templEmbed embed.FS

	dbQuery *chap5.Queries

	store = sessions.NewCookieStore([]byte("forDemo"))
)

// it renders file and push data (d) into template to be rendered
func renderFiles(tmpl string, w http.ResponseWriter, d interface{}) {
	t, err := template.ParseFS(templEmbed, fmt.Sprint("tmpl/%s.html", tmpl))
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}

// securityMiddleWare is a middleware to make sur all requests have valid session and authcated
func securityMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//first of all, all requests MUST have a valid session
		if sessionValid(w, r) {
			//login path will be let throught, otherwise it won't be served to the frontend
			if r.URL.Path == "/login" {
				next.ServeHTTP(w, r)
				return
			}
		}

		// if it does have a new session make sure it has been authenticated
		if hasBeenAuthenticated(w, r) {
			next.ServeHTTP(w, r)
			return
		}

		//otherwise it will need to be redirected to /login
		storeAuthenticated(w, r, false)
		http.Redirect(w, r, "/login", 307)
	})
}

// sessionValid check wheter the session is a valid session
func sessionValid(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session_token")
	return !session.IsNew
}

// authenticationHandler handles authenticaton
func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	result := "Login "
	r.ParseForm()

	if validateUser(r.FormValue("username"), r.FormValue("password")) {
		storeAuthenticated(w, r, true)
		result += "successfull"
	} else {
		result += "UNSUCCESSFULL"
	}
	renderFiles("msg", w, result)
}

// hasBeenAutenticated checks whether the session contain the flag to indicated that
// that the session has gone through authentication process
func hasBeenAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session_token")
	a, _ := session.Values["authenticated"]
	if a == nil {
		return false
	}
	return a.(bool)
}

// storeAuthenticated to store autenticated value
func storeAuthenticated(w http.ResponseWriter, r *http.Request, v bool) {
	session, _ := store.Get(r, "session_token")

	session.Values["authenticated"] = v
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// validateUser check whether username/password exists in database
func validateUser(username, password string) bool {
	//query the data from database
	ctx := context.Background()
	u, _ := dbQuery.GetUserByName(ctx, username)

	//username doesn't exist
	if u.UserName != username {
		return false
	}
	return pkg.CheckPasswordHash(password, u.PassWordHash)
}

func basicMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware called on", r.URL.Path)
		// do stuff
		h.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Server version:", version)
	initDatabase()

	router := mux.NewRouter()

	//POST handler for /login
	router.HandleFunc("/login", authenticationHandler).Methods("POST")

	//embed handler for /css path
	csscontentStatic, _ := fs.Sub(cssEmbed, "css")
	css := http.FileServer(http.FS(csscontentStatic))
	router.PathPrefix("/app").Handler(securityMiddleWare(http.StripPrefix("/css", css)))

	//embed handler for /app path
	contentStatic, _ := fs.Sub(staticEmbed, "static")
	static := http.FileServer(http.FS(contentStatic))
	router.PathPrefix("/app").Handler(securityMiddleWare(http.StripPrefix("/app", static)))

	//add /login path
	router.PathPrefix("/login").Handler(securityMiddleWare(http.StripPrefix("/login", static)))

	//root will redirect to /apo
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusPermanentRedirect)
	})

	//use out basic middleware
	router.Use(basicMiddleWare)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func initDatabase() {
	dbURI := fmt.Sprint("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASS", "mysecretpass"),
		GetAsString("DB_HOST", "localhost"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	// open database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	// connectivity check
	if err := db.Ping(); err != nil {
		log.Fatalln("error pinging db:", err)
	}

	// create the store
	dbQuery = chap5.New(db)

	ctx := context.Background()
	createUserDb(ctx)

	if err != nil {
		os.Exit(1)
	}
}

func createUserDb(ctx context.Context) {
	//has the user been created
	u, _ := dbQuery.GetUserByName(ctx, "user@user")
	if u.UserName == "user@user" {
		log.Println("[!] user@user exists! ...")
		return
	}
	log.Println("[.] creating user: user@user")
	hashPwd, _ := pkg.HashPassword("passsword")
	_, err := dbQuery.CreateUsers(ctx, chap5.CreateUsersParams{
		UserName:     "user@user",
		PassWordHash: hashPwd,
		Name:         "Dummy user",
	})
	if err != nil {
		log.Println("[-] error getting user@dummyuser.domain:", err)
		os.Exit(1)
	}
}
