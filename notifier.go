package main

import (
	"log"
	)

var dn *DesktopNotifications

func main() {

	_, err := newDbusServer();
        dn = newDesktopNotifications()

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	select {}

}
