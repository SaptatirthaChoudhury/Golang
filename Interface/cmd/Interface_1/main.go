package main

import (
	"fmt"
)

type Notifier interface {
	SendNotification(message string)
}

type NotificationGateway struct {
	notify Notifier
}

func (n NotificationGateway) notification(message string) {
	n.notify.SendNotification(message)
}

type EmailNotifier struct {
	Email string
}

func (email EmailNotifier) SendNotification(message string) {
	fmt.Println("Sending Email notification", message)
}

type SMSNotifier struct {
	phone string
}

func (sms SMSNotifier) SendNotification(message string) {
	fmt.Println("Sending sms notification", message)
}

func main() {
	emailNotify := EmailNotifier{Email: "saptatirtha@gmail.com"}
	notify := NotificationGateway{
		notify: emailNotify,
	}
	notify.notification("email:100")

	smsNotify := SMSNotifier{phone: "6001495259"}
	notify = NotificationGateway{
		notify: smsNotify,
	}
    notify.notification("sms:300")
}
