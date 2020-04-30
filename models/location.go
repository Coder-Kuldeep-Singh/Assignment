package models

import (
	"emailsender/db"
	"log"
)

func GetLocation() []string {
	db, err := db.InitDB()
	defer db.Close()
	if err != nil {
		log.Printf("Database Connection String Failed {%v}\n", err)
		return nil
	}
	query, err := db.Query(`SELECT client_user_timezone from emailsender.ClientUserEmail`)
	defer query.Close()
	if err != nil {
		log.Printf("Query Failed! {%v}\n", err)
		return nil
	}
	var UserTimeZone []string
	for query.Next() {
		var usertimezone string
		err = query.Scan(&usertimezone)
		if err != nil {
			log.Println("Error to Execute Statements")
			return nil
		}
		UserTimeZone = append(UserTimeZone, usertimezone)
	}
	return UserTimeZone
}
