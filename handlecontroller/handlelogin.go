package handlecontroller

import (
	"fmt"
	"forum/controller"
	"forum/structure"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// this function handle the login page
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	structure.Tmpl = template.Must(template.ParseFiles("views/login.html"))

	ck, err := r.Cookie("email")

	if r.URL.Path == "/register" {
		HandleRegister(w, r)
		return
	}

	if r.Method != "POST" && err != nil {
		structure.Tmpl.Execute(w, "")
		return
	}
	
	if err != nil {
		email := r.FormValue("email")
		password := r.FormValue("password")
	
		datas, err := controller.GetUser(email)
		if err != nil {
			fmt.Println(err)
			return
		}
	
		if len(datas) == 0 {
			structure.Tmpl.Execute(w, "Email Incorrect !")
			return
		}
	
		structure.User = datas[0]
	
		err = bcrypt.CompareHashAndPassword([]byte(structure.User.Password), []byte(password))
		if err != nil {
			structure.Tmpl.Execute(w, "Password Incorrect !")
			return
		}
	} else {
		email := ck.Value

		datas, err := controller.GetUser(email)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(datas) == 0 {
			structure.Tmpl.Execute(w, "You have an incorrect cookie !")
			return
		}

		structure.User = datas[0]
	}

	cookieEmail := http.Cookie{
		Name:  "email",
		Value: structure.User.Email,
	}

	http.SetCookie(w, &cookieEmail)

	http.Redirect(w, r, "/connected", http.StatusSeeOther)
}
