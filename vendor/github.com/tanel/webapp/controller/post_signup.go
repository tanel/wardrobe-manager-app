package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/model"
	"github.com/tanel/webapp/session"
	"golang.org/x/crypto/bcrypt"
)

// PostSignup creates a new user account
func PostSignup(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := db.SelectUserByEmail(databaseConnection, user.Email, &user); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if user.ID == "" {
		user.ID = uuid.NewV4().String()
		user.CreatedAt = time.Now()

		b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(w, "Password error", http.StatusInternalServerError)
			return
		}

		user.PasswordHash = string(b)

		if err := db.InsertUser(databaseConnection, user); err != nil {
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

	if err := sessionStore.SetUserID(w, r, user.ID); err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
