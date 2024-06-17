package controller

import "forum/structure"

// this function add a user in the table User of the BDD
func AddUser(user structure.UserStruct) (int64, error) {
	db, err := ConnexionToBDD()
	if err != nil {
		return 0, err
	}

	i, err := db.Exec("INSERT INTO `User`(`Username`, `Email`, `Password`) VALUES(?,?,?)", user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	return i.LastInsertId()
}

// This function get the datas of a user in the BDD by this email
func GetUser(email string) ([]structure.UserStruct, error) {
	db, err := ConnexionToBDD()
	if err != nil {
		return []structure.UserStruct{}, err
	}

	datas, err := db.Query("SELECT * FROM `USER` WHERE `Email`=?", email)
	var data []structure.UserStruct

	for datas.Next() {
		var temp structure.UserStruct

		err = datas.Scan(&temp.Id, &temp.Username, &temp.Email, &temp.Password)
		if err != nil {
			return []structure.UserStruct{}, err
		}

		data = append(data, temp)
	}

	return data, nil
}
