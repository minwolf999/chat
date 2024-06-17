package handlecontroller

import (
	"encoding/json"
	"fmt"
	"forum/controller"
	"forum/structure"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// this function handle the chat page
func HandleConnected(w http.ResponseWriter, r *http.Request) {
	structure.Tmpl = template.Must(template.ParseFiles("views/connected.html"))

	ck, err := r.Cookie("email")

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	datas, err := controller.GetUser(ck.Value)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(datas) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	structure.User = datas[0]
	var message string

	if r.Header.Get("Upgrade") == "websocket" {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		structure.Connected[structure.User.Email] = conn

		for {
			_, tempMessage, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}

			message = string(tempMessage)
			fmt.Println("message send to the server !")
			if message != "" {
				err := controller.AddMessage(message, structure.Html)
				if err != nil {
					fmt.Println(err)
					break
				}

				datas, err := controller.GetMessage(structure.Html)
				if err != nil {
					fmt.Println(err)
					break
				}

				userConn, isConnected := structure.Connected[structure.Html.User.Email]
				if isConnected && message != "" {
					jsonMessage, err := json.Marshal(datas)
					if err != nil {
						fmt.Println(err)
						break
					}

					err = userConn.WriteMessage(websocket.TextMessage, jsonMessage)
					if err != nil {
						fmt.Println(err)
						break
					}
				}

				destinataireConn, isConnected := structure.Connected[structure.Html.Destinataire.Email]
				if isConnected && message != "" && structure.Html.Destinataire.Email != structure.Html.User.Email {
					jsonMessage, err := json.Marshal(datas)
					if err != nil {
						fmt.Println(err)
						break
					}

					err = destinataireConn.WriteMessage(websocket.TextMessage, jsonMessage)
					if err != nil {
						fmt.Println(err)
						break
					}
				}
			}
		}
	}

	if structure.User.Email == "" {
		HandleLogin(w, r)
		return
	}

	structure.Html.User = structure.User

	if r.Method != "POST" {
		structure.Tmpl.Execute(w, structure.Html)
		return
	}

	emailDestinataire := r.FormValue("destinataire")
	if emailDestinataire != "" {
		datas, err := controller.GetDestinataire(emailDestinataire)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(datas) == 0 {
			structure.Html.Destinataire.Username = "There is no account using this email address"
		} else {
			structure.Html.Destinataire = datas[0]
		}
	}

	if structure.Html.Destinataire.Username != "" {
		datas, err := controller.GetMessage(structure.Html)
		if err != nil {
			fmt.Println(err)
			return
		}

		structure.Html.Messages = datas
	}

	if r.Header.Get("Upgrade") != "websocket" {
		structure.Tmpl.Execute(w, structure.Html)
	}
}
