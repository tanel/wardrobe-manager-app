package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/form"
	"github.com/tanel/webapp/image"
	"github.com/tanel/webapp/model"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
	"golang.org/x/crypto/bcrypt"
)

// Request represents a HTTP request
type Request struct {
	w  http.ResponseWriter
	r  *http.Request
	ps httprouter.Params
}

const maxUploadSize = 5 * 1024 * 1024

// NewRequest returns new request instance
func NewRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*Request, error) {
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
		w:  w,
		r:  r,
		ps: ps,
	}, nil
}

// Unmarshal unmarshals JSON
func (request *Request) Unmarshal(v interface{}) bool {
	b, err := ioutil.ReadAll(request.r.Body)
	if err != nil {
		log.Println(err)
		http.Error(request.w, "reading request body failed", http.StatusInternalServerError)
		return false
	}

	defer func() {
		if closeErr := request.r.Body.Close(); closeErr != nil {
			log.Println(errors.Annotate(closeErr, "closing request body failed"))
		}
	}()

	if err := json.Unmarshal(b, v); err != nil {
		request.BadRequest(err.Error())
		return false
	}

	return true
}

// Marshal marshals JSON
func (request *Request) Marshal(v interface{}) bool {
	b, err := json.Marshal(v)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "marshalling JSON failed"))
		return false
	}

	return request.Write(b)
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
	value, err := session.SharedInstance.Value(request.r, key)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting value from session failed"))
		return nil, false
	}

	return value, true
}

// SetSessionValue sets value in session
func (request *Request) SetSessionValue(key, value string) bool {
	if err := session.SharedInstance.SetValue(request.w, request.r, key, value); err != nil {
		request.InternalServerError(errors.Annotate(err, "setting value in session failed"))
		return false
	}

	return true
}

// CurrentUser returns logged in user
func (request *Request) CurrentUser() (*model.User, bool) {
	userID, err := session.SharedInstance.UserID(request.r)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting user ID from session failed"))
		return nil, false
	}

	user, err := db.SelectUserByID(*userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting user by ID failed"))
		return nil, false
	}

	return user, true
}

// UserID returns logged in user ID
func (request *Request) UserID() (*string, bool) {
	userID, err := session.SharedInstance.UserID(request.r)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting user ID from session failed"))
		return nil, false
	}

	return userID, true
}

// SetUserID logs user in
func (request *Request) SetUserID(userID string) bool {
	if err := session.SharedInstance.SetUserID(request.w, request.r, userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "setting user ID in session failed"))
		return false
	}

	return true
}

// ClearUserID logs user out
func (request *Request) ClearUserID() bool {
	if err := session.SharedInstance.SetUserID(request.w, request.r, ""); err != nil {
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
	http.Error(request.w, message, http.StatusUnauthorized)
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
func (request *Request) File(name string) ([]byte, bool) {
	b, err := form.File(request.r, name)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting file upload failed"))
		return nil, false
	}

	return b, true
}

// SaveFormImage saves image uploaded with HTML form
func (request *Request) SaveFormImage(name, userID string) bool {
	b, ok := request.File(name)
	if !ok {
		return false
	}

	if len(b) == 0 {
		return true
	}

	img := model.Image{
		Base: model.Base{
			ID:        uuid.Must(uuid.NewV4()).String(),
			CreatedAt: time.Now(),
		},
		UserID: userID,
		Body:   b,
	}
	if err := db.InsertImage(img); err != nil {
		request.InternalServerError(errors.Annotate(err, "inserting image failed"))
		return false
	}

	if err := img.Save(); err != nil {
		request.InternalServerError(errors.Annotate(err, "saving upladed file to disk failed"))
		return false
	}

	if err := image.GenerateThumbnailsForImage(img.FilePath()); err != nil {
		request.InternalServerError(errors.Annotate(err, "getting file upload failed"))
		return false

	}

	return true
}

// SaveImage saves image
func (request *Request) SaveImage(b []byte, name, userID string) bool {
	img := model.Image{
		Base: model.Base{
			ID:        uuid.Must(uuid.NewV4()).String(),
			CreatedAt: time.Now(),
		},
		UserID: userID,
		Body:   b,
	}
	if err := db.InsertImage(img); err != nil {
		request.InternalServerError(errors.Annotate(err, "inserting image failed"))
		return false
	}

	if err := img.Save(); err != nil {
		request.InternalServerError(errors.Annotate(err, "saving upladed file to disk failed"))
		return false
	}

	if err := image.GenerateThumbnailsForImage(img.FilePath()); err != nil {
		request.InternalServerError(errors.Annotate(err, "getting file upload failed"))
		return false

	}

	return true
}

// CreateUser creates user
func (request *Request) CreateUser(email, password string) (*model.User, bool) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "hashing password failed"))
		return nil, false
	}

	passwordHash := string(b)
	user := model.User{
		Base: model.Base{
			ID:        uuid.Must(uuid.NewV4()).String(),
			CreatedAt: time.Now(),
		},
		Name:         &email,
		Email:        email,
		PasswordHash: &passwordHash,
	}
	if err := db.InsertUser(user); err != nil {
		request.InternalServerError(errors.Annotate(err, "inserting user into database failed"))
		return nil, false
	}

	return &user, true

}
