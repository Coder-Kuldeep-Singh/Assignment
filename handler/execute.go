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
func Execute(body string, usertimezone []string, wg *sync.WaitGroup) {
	var AllLocations string

	for _, Location := range usertimezone {
		wg.Add(1)

		loc, err := time.LoadLocation(Location)
		if err != nil {
			log.Printf("Error In Location {%v}\n", err)
		}

		current_time := time.Now().In(loc)

		destinationYear := current_time.Year()
		destinationMonth := current_time.Month()
		destinationDay := current_time.Day()
		destinationHour := current_time.Hour()
		destinationMinute := current_time.Minute()
		destinationSeconds := current_time.Second()

		LocationTime := fmt.Sprintf("%v-%v-%v %v:%v:%v", destinationYear, destinationMonth, destinationDay, destinationHour, destinationMinute, destinationSeconds)
		fmt.Printf("Location: {%v} LocationTime: {%v}\n", Location, LocationTime)

		// Time to match which time we have to send the email to clients
		Time := fmt.Sprintf("%v-%v-%v 18:0:0", destinationYear, destinationMonth, destinationDay)
		fmt.Println(Time)

		Comparision := Time

		if LocationTime == Comparision {
			fmt.Println(Location)
			AllLocations = Location
		}
	}
	fmt.Println("*************************************************************************************************************")

	user := models.GetEmails(AllLocations)
	if user == nil {
		return
	}

	// Send Email
	for _, users := range user {
		wg.Add(1)
		go mails.SendEmail(body, users, wg)
	}

	wg.Wait()
}
