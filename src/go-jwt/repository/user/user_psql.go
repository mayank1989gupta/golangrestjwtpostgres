package userrepository

import (
	"database/sql"
	"log"

	"../../models"
)

//UserRepository - User Repository Struct
type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Signup - Signup method - interaction with DB
func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	//inserting the record into DB
	query := "insert into users (email, password) values ($1, $2) RETURNING id;"
	//Scan is used as we expect the query to return ID -> scan returns err or, nil
	err := db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	logFatal(err)
	return user
}

//Login - Method to ineract wit db for login
func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	//QueryRow -> returns atmost 1 row
	row := db.QueryRow("select * from users where email = $1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password) //scan maps the data to struct

	if err != nil {
		return user, err
	}

	return user, nil
}
