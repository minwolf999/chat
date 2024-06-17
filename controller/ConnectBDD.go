package controller

import "database/sql"

// open a connection with the DBB
func ConnexionToBDD() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "structure/database.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
