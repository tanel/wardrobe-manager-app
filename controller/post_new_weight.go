package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/session"
)

// PostNewWeight saves a new weight into database
func PostNewWeight(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	weightEntry, err := model.NewWeightEntryForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weightEntry.ID = uuid.NewV4().String()
	weightEntry.UserID = userID

	if err := db.InsertWeight(databaseConnection, *weightEntry); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/weight", http.StatusSeeOther)
}
