package render

import (
	"bytes"
	"github.com/dapetoo/go-bookings/internal/config"
	"github.com/dapetoo/go-bookings/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		//Get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//Get the requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//Create a map to act as the cache
	myCache := map[string]*template.Template{}

	//Get all the page files
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//Loop through the pages one-by-one
	for _, page := range pages {
		//Get the file name
		name := filepath.Base(page)

		//Parse the page template file in to the template set
		ts, err := template.ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//Parse the template set to the cache
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		//If there is a match, parse the template file in to the template set
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		//Add the template set to the cache
		myCache[name] = ts
	}
	return myCache, nil
}

//func RenderTemplate(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		log.Println("Error parsing template:", err)
//		return
//	}
//}
//
//var tc = make(map[string]*template.Template)
//
//func RenderTemplateTest(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	// Check if the template cache exists
//	_, inMap := tc[t]
//	if !inMap {
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println("error creating template cache:", err)
//			// Handle the error appropriately
//			return
//		}
//		log.Println("template not in cache")
//	} else {
//		log.Println("using cached template")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println("error executing template:", err)
//	}
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//
//	//Parse the template
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//
//	//Add the template to the cache
//	tc[t] = tmpl
//	return nil
//}
