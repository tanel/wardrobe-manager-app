package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"github.com/tanel/wardrobe-manager-app/session"
	"golang.org/x/crypto/bcrypt"
)

func GetSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := session.UserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID != nil {
		http.Redirect(w, r, frontpage, http.StatusSeeOther)
		return
	}

	if err := Render(w, "signup", nil); err != nil {
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

	if err := db.SelectUserByEmail(user.Email, &user); err != nil {
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

		if err := db.InsertUser(user); err != nil {
			log.Println(err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
			log.Println(err)
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
	}

	if err := session.SetUserID(w, r, user.ID); err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, frontpage, http.StatusSeeOther)
}
