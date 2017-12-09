package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/model"
	"golang.org/x/crypto/bcrypt"
)

func Serve(port string) {
	router := httprouter.New()

	router.GET("/signup", GetSignup)
	router.POST("/signup", PostSignup)
	router.GET("/", GetIndex)

	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	log.Println("Server starting at http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, router))
}

func GetSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := filepath.Join("templates", "*")
	list, err := filepath.Glob(path)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.New("signup").ParseFiles(list...)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	password := r.FormValue("password")

	var user model.User
	user.Email = r.FormValue("email")

	err := db.QueryRow("SELECT id, password_hash FROM users WHERE email = $1", user.Email).Scan(&user.ID, &user.PasswordHash)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID == "" {
		user.ID = uuid.NewV4().String()

		b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user.PasswordHash = string(b)

		_, err = db.Exec("INSERT INTO users(id, email, password_hash) VALUES($1, $2, $3)", user.ID, user.Email, user.PasswordHash)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}

	// create session
	// redirect
}

func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := filepath.Join("templates", "*")
	list, err := filepath.Glob(path)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.New("index").ParseFiles(list...)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
