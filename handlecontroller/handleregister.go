package handlecontroller

import (
	"forum/controller"
	"forum/structure"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// this function handle the register page
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	structure.Tmpl = template.Must(template.ParseFiles("views/register.html"))

	if r.URL.Path == "/login" {
		HandleLogin(w, r)
		return
	}

	if r.Method != "POST" {
		structure.Tmpl.Execute(w, "")
		return
	}

	structure.User.Username = r.FormValue("username")
	structure.User.Email = r.FormValue("email")
	structure.User.Password = r.FormValue("password")

	hashed, err := bcrypt.GenerateFromPassword([]byte(structure.User.Password), 14)
	if err != nil {
		structure.Tmpl.Execute(w, "Wrong password")
		return
	}

	structure.User.Password = string(hashed)

	id, err := controller.AddUser(structure.User)
	if err != nil {
		structure.Tmpl.Execute(w, "Email Already used")
		return
	}

	structure.User.Id = int(id)

	cookieEmail := http.Cookie{
		Name:  "email",
		Value: structure.User.Email,
	}

	http.SetCookie(w, &cookieEmail)

	http.Redirect(w, r, "/connected", http.StatusSeeOther)
}
