package main

import (
	"log"
        "github.com/godbus/dbus"
)

const busName = "com.NSAway.EventNotifier"
const objectPath = "/com/NSAway/EventNotifier"

type dbusObject struct {
        dbus.BusObject
}

type slmData struct {
    Msg string
    FullScreen bool
    Priority int
}

func newDbusObject() (*dbusObject, error) {
        conn, err := dbus.SystemBus()

        if err != nil {
                return nil, err
        }

        return &dbusObject{conn.Object(busName, objectPath)}, nil
}

func (ob *dbusObject) alertObj(Msg string, Priority int) {
    dobj := slmData{Msg, Priority}
    call := ob.Call("com.NSAway.EventNotifier.Alert", 0, dobj)
    if call.Err != nil {
            log.Println("errore")
	    panic(call.Err)
    }
}


