package main

import (
	"fmt"
	"log"

	"github.com/godbus/dbus"
	"github.com/TheCreeper/go-notify"
)

const (
	INFO = 0
	NOTICE = 1
	WARNING = 2
	CRITICAL = 3
)



type DesktopNotifications struct {
	notifications map[string]uint32
}

func newDesktopNotifications() *DesktopNotifications {
	if _, err := dbus.SessionBus(); err != nil {
		log.Printf("Error enabling dbus based notifications! %+v\n", err)
		return nil
	}

	dn := new(DesktopNotifications)
	dn.notifications = make(map[string]uint32)
	return dn
}

func (dn *DesktopNotifications) show(cat, message string,  Priority int) error {

	hints := make(map[string]interface{})
	hints[notify.HintTransient] = false
	hints[notify.HintActionIcons] = "NSAway"
	icon := "dialog-warning"
	timeout := notify.ExpiresDefault
	showFullscreen := false

	if Priority == NOTICE {
		icon = "dialog-information"
	}

	if Priority == WARNING {
		timeout = 15000
	}
	if Priority == CRITICAL {
		timeout = notify.ExpiresNever
		icon = "dialog-error"
		showFullscreen = true
	}

	if showFullscreen {
		hints[notify.HintUrgency] = notify.UrgencyCritical
	}


	notification := notify.Notification{
		AppName: "EventNotifier",
		AppIcon: icon,
		Timeout: int32(timeout),
		Hints:   hints,
	}
	notification.Summary = "NSAway Event Notifier"
	notification.Body = message

	nid, err := notification.Show()
	if err != nil {
		return fmt.Errorf("Error showing notification: %v", err)
	}
	dn.notifications[cat] = nid
	return nil
}

