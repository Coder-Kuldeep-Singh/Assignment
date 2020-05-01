package models

import (
	"emailsender/db"
	"fmt"
	"log"
)

// GetEmails function Gets the emails of user ana client from database
func GetEmails(Location string) []string {
	fmt.Println(Location)
	db, err := db.InitDB()
	defer db.Close()
	if err != nil {
		log.Printf("Database Connection String Failed {%v}\n", err)
		return nil
	}
	query, err := db.Query("SELECT Client_user_email from emailsender.ClientUserEmail where client_user_timezone = " + "'" + Location + "'" + ";")
	defer query.Close()
	if err != nil {
		log.Printf("Email Query Failed! {%v}\n", err)
		return nil
	}
	var UserEmail []string
	for query.Next() {
		var user string
		err = query.Scan(&user)
		if err != nil {
			log.Println("Error in Second Query {%v}", err)
			return nil
		}
		UserEmail = append(UserEmail, user)
	}
	return UserEmail
}
