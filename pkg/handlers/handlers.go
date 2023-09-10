package handlers

import (
	"net/http"

	"github.com/kpauletti/gwa/pkg/config"
	"github.com/kpauletti/gwa/pkg/models"
	"github.com/kpauletti/gwa/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.App.SessionManager.GetString(r.Context(), "remote_ip")

	stringMap := map[string]string{}
	stringMap["test"] = remoteIP + " is the remote IP"

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
