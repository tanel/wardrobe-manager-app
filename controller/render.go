package controller

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/juju/errors"
)

// Render renders a template with given data
func Render(w io.Writer, templateName string, data interface{}) error {
	path := filepath.Join("template", "*.html")
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
