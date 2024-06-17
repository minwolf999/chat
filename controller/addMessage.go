package controller

import (
	"forum/structure"
	"time"
)

// Add a message in the table Message of the BDD
func AddMessage(message string, html structure.Htmls) error {
	db, err := ConnexionToBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO `Message`(`Date`,`Id_Sender`,`Id_Receiver`,`Message`) VALUES(?,?,?,?)", time.Now().Format("02/01/2006 15:04:05"), html.User.Id, html.Destinataire.Id, message)
	return err
}
