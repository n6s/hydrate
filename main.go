package main

import (
	// "fmt"
	"github.com/0xAX/notificator"
	"time"
)

var notify *notificator.Notificator

func main() {
	title := "Hydrated?"
	message := "Have another drink!"
	iconpath := "/home/roger/Pictures/water.png"
	for {
		// fmt.Println(time.Now().Format("15:04"))
		// fmt.Printf("%s\n", message)
		notifyMe(title, message, iconpath)

		timer1 := time.NewTimer(time.Minute * 5)
		<-timer1.C
	}
}

func notifyMe(title, message, icon string) {

	notify = notificator.New(notificator.Options{
		DefaultIcon: "/usr/share/icons/elementary-xfce/notifications/24/notification-message-im.png",
		AppName:     "My test App",
	})

	notify.Push(title, message, icon, notificator.UR_CRITICAL)
}
