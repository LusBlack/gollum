package render

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/LusBlack/gollum/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}

	// execute and save
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, nil); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}

	//write to browser
	_, err := buf.WriteTo(w)
	if err != nil {
		http.Error(w, "Error writing to browser", http.StatusInternalServerError)
	}
}

// creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}

/*
parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
err = parsedTemplate.Execute(w, nil)
if err != nil {
	fmt.Println("Error parsing template", err)
	return
}
*/
