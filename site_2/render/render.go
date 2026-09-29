package render

import (
	"app/pkg/config"
	tmplData "app/pkg/templateData"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}


func RenderTmplCache(w http.ResponseWriter, name string, data *tmplData.TemplateData) {
	var tmpls map[string]*template.Template

	if app.UseCache {
		tmpls = app.TemplatesCache
	} else {
		tmpls, _ = CreateTemplatesCache()
	}

	tmpl, ok := tmpls[name]
	if !ok {
		log.Println("Haven't find tmpl")
	}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)

	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplatesCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../tmpls/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)
		tmpl, err := template.New(pageName).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("../tmpls/*.layout.tmpl")

		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob("../tmpls/*.layout.tmpl")

			if err != nil {
				return cache, err
			}
		}

		cache[pageName] = tmpl
	}

	return cache, nil
}
