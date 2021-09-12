package render

import (
	"html/template"
	"net/http"
	"path"

	"github.com/catehulu/factorio-calculator/internal/models"
)

/** RenderTemplate render template using HTML template */
func RenderTemplate(rw http.ResponseWriter, vpath string, td *models.TemplateData) {

	var filepath = path.Join("D:/Works/Mine/factorio-calculator", "templates", vpath)
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(rw, td)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
