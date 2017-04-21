package main
import (
//	"fmt"
	"log"
)

func main(){
	dbo, err := newDbusObject()
	if err != nil {
		log.Fatal("Error connecting to SystemBus: %v", err)
	}

	// obj.alertObj(message,showFullscreen,Priority)
	//Priority has 4 values, 0 for NOTICE, 1 for INFO 2 for WARNING  and 3 for CRITICAL

	dbo.alertObj("INFO",1)

	//dbo.alertObj("WARNING",2)

	//dbo.alertObj("CRITICAL",3)
}
