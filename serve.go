package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"golang.org/x/crypto/bcrypt"
)

func Serve(port string) {
	router := httprouter.New()

	router.GET("/signup", GetSignup)
	router.POST("/signup", PostSignup)
	router.GET("/wardrobe", GetWardrobe)
	router.GET("/logout", GetLogout)
	router.GET("/", GetIndex)

	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	log.Println("Server starting at http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, router))
}

func GetSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := sessionUserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID != nil {
		http.Redirect(w, r, "/wardrobe", http.StatusSeeOther)
		return
	}

	path := filepath.Join("templates", "*")
	list, err := filepath.Glob(path)
	if err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}

	t, err := template.New("signup").ParseFiles(list...)
	if err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func PostSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "form error", http.StatusInternalServerError)
		return
	}

	var user model.User
	user.Email = strings.TrimSpace(r.FormValue("email"))
	if user.Email == "" {
		http.Error(w, "please enter an e-mail", http.StatusBadRequest)
		return
	}

	password := strings.TrimSpace(r.FormValue("password"))
	if password == "" {
		http.Error(w, "please enter a password", http.StatusBadRequest)
		return
	}

	err := db.SelectUserByEmail(user.Email, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if user.ID == "" {
		user.ID = uuid.NewV4().String()

		b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(w, "Password error", http.StatusInternalServerError)
			return
		}

		user.PasswordHash = string(b)

		err = db.InsertUser(user)
		if err != nil {
			log.Println(err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
	}

	if err := setSessionUserID(w, r, user.ID); err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/wardrobe", http.StatusSeeOther)
}

func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := sessionUserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID != nil {
		http.Redirect(w, r, "/wardrobe", http.StatusSeeOther)
		return
	}

	path := filepath.Join("templates", "*")
	list, err := filepath.Glob(path)
	if err != nil {
		log.Println(err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	t, err := template.New("index").ParseFiles(list...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}

func GetLogout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := setSessionUserID(w, r, ""); err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetWardrobe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := filepath.Join("templates", "*")
	list, err := filepath.Glob(path)
	if err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}

	t, err := template.New("wardrobe").ParseFiles(list...)
	if err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
