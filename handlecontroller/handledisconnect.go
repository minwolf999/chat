package handlecontroller

import (
	"forum/structure"
	"net/http"
)

// this function handle the disconnect route
func HandleDisconnect(w http.ResponseWriter, r *http.Request) {
	delete(structure.Connected, structure.Html.User.Email)
}
