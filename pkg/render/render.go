package render

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/kpauletti/gwa/pkg/config"
	"github.com/kpauletti/gwa/pkg/models"
)

// template cache
var tc = map[string]*template.Template{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// renders a template from cache or file
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc = map[string]*template.Template{}

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//confirm template exists
	t, ok := tc[tmpl]

	if !ok {
		app.ErrorLog.Fatal("Could not get template from template cache")
	}

	//execute it to a buffer first, just an extra step to ensure template is working
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	//write it to the http response
	_, err = buf.WriteTo(w)

	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
