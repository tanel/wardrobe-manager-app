package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"golang.org/x/crypto/bcrypt"
)

const sessionName = "wardrobe-app-session"

// FIXME: get secret from environment
var store = sessions.NewCookieStore([]byte("C93B74DA-4D85-418C-B513-3BEDE6BFCECC"))

func Serve(port string) {
	router := httprouter.New()

	router.GET("/signup", GetSignup)
	router.POST("/signup", PostSignup)
	router.GET("/logout", GetSignup)
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

	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["user_id"] = user.ID

	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/wardrobe", http.StatusSeeOther)
}

func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["user_id"] = nil

	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
