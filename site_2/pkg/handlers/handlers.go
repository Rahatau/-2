package handlers

import (
	"net/http"
	"app/pkg/templateData"
	"app/pkg/config"
	"app/render"
)


type Repository struct {
	App config.AppConfig
}

var Repo *Repository

func NewRepo(app *config.AppConfig) *Repository {
	Repo = &Repository{
		App: *app,
	}

	return Repo
}


func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTmplCache(w, "home.page.tmpl", &tmplData.TemplateData{StringMap: map[string]string{"word": "Hello1"}})

	
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTmplCache(w, "about.page.tmpl", &tmplData.TemplateData{StringMap: map[string]string{"word": "Hello2"}})
}
