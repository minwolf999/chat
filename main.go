package main

import (
	"fmt"
	"forum/handlecontroller"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// createBDD
}

func main() {
	fmt.Println("starting server at port :8080")
	fmt.Println("http://localhost:8080")

	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	http.HandleFunc("/", handlecontroller.HandleLogin)
	http.HandleFunc("/connected", handlecontroller.HandleConnected)
	http.HandleFunc("/profile", handlecontroller.HandleProfile)
	http.HandleFunc("/bye", handlecontroller.HandleDisconnect)
	http.ListenAndServe(":8080", nil)

}
