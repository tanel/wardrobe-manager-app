package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
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

	email := r.FormValue("email")
	password := r.FormValue("password")

	// search from database by email
	// 		if does not exists, create new user
	// 		if exists, compare password
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
