package http

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/form"
	"github.com/tanel/webapp/model"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
)

// Request represents a HTTP request
type Request struct {
	DB           *sql.DB
	sessionStore *session.Store
	w            http.ResponseWriter
	r            *http.Request
	ps           httprouter.Params
}

const maxUploadSize = 5 * 1024 * 1024

// NewRequest returns new request instance
func NewRequest(db *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*Request, error) {
	var err error
	if r.Header.Get("Content-Type") == "multipart/form-data" {
		err = r.ParseMultipartForm(maxUploadSize)
	} else {
		err = r.ParseForm()
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "form error", http.StatusInternalServerError)
		return nil, errors.Annotate(err, "parsing form failed")
	}

	return &Request{
		DB:           db,
		sessionStore: sessionStore,
		w:            w,
		r:            r,
		ps:           ps,
	}, nil
}

// Query returns URL query
func (request *Request) Query() url.Values {
	return request.r.URL.Query()
}

// R returns r
func (request *Request) R() *http.Request {
	return request.r
}

// SessionValue returns value from session
func (request *Request) SessionValue(key string) (*string, bool) {
	value, err := request.sessionStore.Value(request.r, key)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting value from session failed"))
		return nil, false
	}

	return value, true
}

// SetSessionValue sets value in session
func (request *Request) SetSessionValue(key, value string) bool {
	if err := request.sessionStore.SetValue(request.w, request.r, key, value); err != nil {
		request.InternalServerError(errors.Annotate(err, "setting value in session failed"))
		return false
	}

	return true
}

// CurrentUser returns logged in user
func (request *Request) CurrentUser() (*model.User, bool) {
	userID, err := request.sessionStore.UserID(request.r)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting user ID from session failed"))
		return nil, false
	}

	user, err := db.SelectUserByID(request.DB, *userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting user by ID failed"))
		return nil, false
	}

	return user, true
}

// UserID returns logged in user ID
func (request *Request) UserID() (*string, bool) {
	userID, err := request.sessionStore.UserID(request.r)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting user ID from session failed"))
		return nil, false
	}

	return userID, true
}

// SetUserID logs user in
func (request *Request) SetUserID(userID string) bool {
	if err := request.sessionStore.SetUserID(request.w, request.r, userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "setting user ID in session failed"))
		return false
	}

	return true
}

// ClearUserID logs user out
func (request *Request) ClearUserID() bool {
	if err := request.sessionStore.SetUserID(request.w, request.r, ""); err != nil {
		request.InternalServerError(errors.Annotate(err, "setting user ID in session failed"))
		return false
	}

	return true
}

// Render renders a template
func (request *Request) Render(templateName string, data interface{}) bool {
	if err := template.Render(request.w, templateName, data); err != nil {
		request.InternalServerError(errors.Annotatef(err, "rendering template %s failed", templateName))
		return false
	}

	return true
}

// Redirect redirects
func (request *Request) Redirect(path string) {
	http.Redirect(request.w, request.r, path, http.StatusSeeOther)
}

// InternalServerError returns internal server error
func (request *Request) InternalServerError(err error) {
	log.Println(err)
	http.Error(request.w, "An error occurred. We're are sorry. Please try again later", http.StatusInternalServerError)
}

// BadRequest returns bad request
func (request *Request) BadRequest(message string) {
	http.Error(request.w, message, http.StatusBadRequest)
}

// Unauthorized returns unauthorized
func (request *Request) Unauthorized(message string) {
	http.Error(request.w, "Invalid password", http.StatusUnauthorized)
}

// FormValue returns form value
func (request *Request) FormValue(key string) string {
	return strings.TrimSpace(request.r.FormValue(key))
}

// ParamByName returns param by name
func (request *Request) ParamByName(name string) string {
	return request.ps.ByName(name)
}

// QueryParamByName returns query param by name
func (request *Request) QueryParamByName(name string) string {
	return request.r.URL.Query().Get(name)
}

// NotFound returns not found
func (request *Request) NotFound(message string) {
	http.Error(request.w, message, http.StatusNotFound)
}

// Write writes bytes
func (request *Request) Write(b []byte) bool {
	if _, err := request.w.Write(b); err != nil {
		log.Println(err)
		return false
	}

	return true
}

// File gets uploaded file from request
func (request *Request) File(name string) ([]byte, error) {
	b, err := form.File(request.r, name)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting file upload failed"))
		return nil, err
	}

	return b, nil
}
