package handlecontroller

import (
	"forum/structure"
	"net/http"
	"text/template"
)

// this function handle the profile page
func HandleProfile(w http.ResponseWriter, r *http.Request) {
	structure.Tmpl = template.Must(template.ParseFiles("views/profile.html"))

	structure.Tmpl.Execute(w, structure.Html.User)
}
