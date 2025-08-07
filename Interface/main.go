package main

import (
	"fmt"
)

// Define Notifier as interface
type Notifier interface {
	SendNotification(message string)
}

// Define EmailNotifier as Struct
type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) SendNotification(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.Email, message)
}

// Define SMSNotifier as struct
type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) SendNotification(message string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.Phone, message)
}

// Function that works with any Notifier
func NotifyUser(n Notifier) {
	n.SendNotification("Your order has been shipped!")
}

func main() {
	email := EmailNotifier{
		Email: "saptatirtha@gmail.com",
	}
	sms := SMSNotifier{
		Phone: "+916001495259",
	}

	NotifyUser(email)
	NotifyUser(sms)
}
