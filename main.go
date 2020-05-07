package main

import (
	"emailsender/handler"
	"emailsender/models"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	usertimezone := models.GetLocation()
	if usertimezone == nil {
		return
	}
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				go handler.Execute("Good Morning", usertimezone, &wg)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println("Finished Execution")
}
