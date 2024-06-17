package controller

import "forum/structure"

// This function get the datas of a user in the BDD by this email
func GetDestinataire(email string) ([]structure.DestinataireStruct, error) {
	db, err := ConnexionToBDD()
	if err != nil {
		return []structure.DestinataireStruct{}, err
	}

	datas, err := db.Query("SELECT * FROM `USER` WHERE `Email`=?", email)
	var data []structure.DestinataireStruct

	for datas.Next() {
		var temp structure.DestinataireStruct

		err = datas.Scan(&temp.Id, &temp.Username, &temp.Email, &temp.Password)
		if err != nil {
			return []structure.DestinataireStruct{}, err
		}

		data = append(data, temp)
	}

	return data, nil
}

// This function get the datas of all the messages in the BDD
func GetMessage(html structure.Htmls) ([]structure.Message, error) {
	db, err := ConnexionToBDD()
	if err != nil {
		return []structure.Message{}, err
	}

	datas, err := db.Query("SELECT m.`Id`, m.`Date`, u.`Username`, u2.`Username`, m.`Message` FROM `Message` m INNER JOIN `User` u ON u.`Id`= m.`Id_Sender` INNER JOIN `User` u2 ON u2.`Id`= m.`Id_Receiver` WHERE (m.`Id_Sender`=? AND m.`Id_Receiver`=?) OR (m.`Id_Sender`=? AND m.`Id_Receiver`=?)", html.User.Id, html.Destinataire.Id, html.Destinataire.Id, html.User.Id)
	if err != nil {
		return []structure.Message{}, err
	}

	var data []structure.Message
	for datas.Next() {
		var temp structure.Message
		err = datas.Scan(&temp.Id, &temp.Date, &temp.Sender, &temp.Receiver, &temp.Message)
		if err != nil {
			return []structure.Message{}, err
		}

		data = append(data, temp)
	}

	return data, nil
}
