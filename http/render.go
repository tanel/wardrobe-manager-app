package http

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/juju/errors"
)

func Render(w http.ResponseWriter, templateName string, data interface{}) error {
	path := filepath.Join("templates", "*")
	list, err := filepath.Glob(path)
	if err != nil {
		return errors.Annotate(err, "globbing templates failed")
	}

	t, err := template.ParseFiles(list...)
	if err != nil {
		return errors.Annotate(err, "parsing templates failed")
	}

	if err := t.ExecuteTemplate(w, templateName, data); err != nil {
		return errors.Annotate(err, "executing template failed")
	}

	return nil
}
