package main

import (
	"log"
	"time"
)

func main() {

	log.Println("Starting IPX800 Watchdogs")

	// based on https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			// do stuff
			go Get_IPX()
		case <-quit:
			ticker.Stop()
			return
		}
	}

	log.Println("Bye bye!")
}
