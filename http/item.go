package http

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"github.com/tanel/wardrobe-manager-app/session"
)

func GetItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := session.UserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	page := ItemsPage{
		Page: Page{
			UserID: *userID,
		},
		Items: nil,
	}
	if err := Render(w, "items", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func GetItemsNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := session.UserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	page := ItemPage{
		Page: Page{
			UserID: *userID,
		},
	}
	if err := Render(w, "items-new", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func PostItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "form error", http.StatusInternalServerError)
		return
	}

	var item model.Item

	item.ID = uuid.NewV4().String()

	item.Name = strings.TrimSpace(r.FormValue("name"))
	if item.Name == "" {
		http.Error(w, "please enter a name", http.StatusBadRequest)
		return
	}

	item.Description = strings.TrimSpace(r.FormValue("description"))
	item.Color = strings.TrimSpace(r.FormValue("color"))
	item.Size = strings.TrimSpace(r.FormValue("size"))
	item.Brand = strings.TrimSpace(r.FormValue("brand"))
	item.Category = strings.TrimSpace(r.FormValue("category"))
	item.Currency = strings.TrimSpace(r.FormValue("currency"))

	s := strings.TrimSpace(r.FormValue("price"))
	if s != "" {
		var err error
		item.Price, err = strconv.ParseFloat(s, 64)
		if err != nil {
			http.Error(w, "please enter a valid price or leave it blank", http.StatusBadRequest)
			return
		}
	}

	if err := db.InsertItem(item); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, frontpage, http.StatusSeeOther)
}
