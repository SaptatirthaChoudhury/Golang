package main

import (
	"fmt"
)

type EmailNotifier struct {
	Email string
}

// Method with value receiver
func (e EmailNotifier) SendNotification(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.Email, message)
}



func Notifier() {
  //Create an instance of EmailNotifier
  notifier := EmailNotifier{Email: "sapta@protonmail.com"} 

  notifier.SendNotification("Just testing")

}
