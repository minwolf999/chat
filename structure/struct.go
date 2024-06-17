package structure

import (
	"text/template"

	"github.com/gorilla/websocket"
)

var (
	Tmpl      *template.Template
	User      UserStruct
	Html      Htmls
	Connected = map[string]*websocket.Conn{}
)

type Htmls struct {
	User         UserStruct
	Destinataire DestinataireStruct
	Messages     []Message
}

type UserStruct struct {
	Id       int
	Username string
	Email    string
	Password string
}

type DestinataireStruct struct {
	Id       int
	Username string
	Email    string
	Password string
}

type Message struct {
	Id       int
	Date     string
	Sender   string
	Receiver string
	Message  string
}
