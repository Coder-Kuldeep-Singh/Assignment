package handler

import (
	mails "emailsender/email"
	"emailsender/models"
	"fmt"
	"log"
	"sync"
	"time"
)

// Execute function Is the combination of the all functions to run
func Execute(body string, wg *sync.WaitGroup) {
	usertimezone := models.GetLocation()
	if usertimezone == nil {
		return
	}
	var AllLocations []string

	for _, Location := range usertimezone {
		wg.Add(1)

		loc, err := time.LoadLocation(Location)
		if err != nil {
			log.Printf("Error In Location {%v}\n", err)
		}

		current_time := time.Now().In(loc)
		destinationTime := current_time.Format("2006-April-02 15:04:05")
		LocationTime := destinationTime
		fmt.Printf("Location: {%v} LocationTime: {%v}\n", Location, LocationTime)

		destinationYear := current_time.Year()
		destinationMonth := current_time.Month()
		destinationDay := current_time.Day()

		// Time to match which time we have to send the email to clients
		Time := fmt.Sprintf("%v-%v-%v 08:00:00", destinationYear, destinationMonth, destinationDay)

		Comparision := Time

		if LocationTime == Comparision {
			fmt.Println(Location)
			AllLocations = append(AllLocations, Location)
		}
	}
	fmt.Println("*************************************************************************************************************")
	for _, Location := range AllLocations {
		user := models.GetEmails(Location)
		if user == nil {
			return
		}
		for _, user := range user {
			wg.Add(1)
			go mails.SendEmail(body, user, wg)
		}
	}
	wg.Wait()
}
