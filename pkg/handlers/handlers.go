package handlers

import (
	"net/http"

	"github.com/LusBlack/gollum/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}
<<<<<<< HEAD

// func RenderTemplateTest(w http.ResponseWriter, tmpl string) error {
// 	myCache := map[string]*template.Template{}

// 	pages, err := filepath.Glob("./templates/*.page.tmpl")
// 	if err != nil {
// 		return err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)

// 		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
// 	}
// }
=======
>>>>>>> 637c688 (fixing git issues after system crash)
