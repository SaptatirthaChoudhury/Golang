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

type PriceCalculator interface {
	// We don't care about parameter names in the interface
	CalculatePrice(string, float64, float64) float64
}

type Product struct{}

func (p Product) CalculatePrice(string, float64, float64) float64 {
	fmt.Println("Parameters ignored, using flat price")
	return 100.0
}

func main() {
	email := EmailNotifier{
		Email: "saptatirtha@gmail.com",
	}
	sms := SMSNotifier{
		Phone: "+916001495259",
	}

	//email.SendNotification("sending email")
	//sms.SendNotification("sending sms")

	NotifyUser(email)
	NotifyUser(sms)

	var calc PriceCalculator = Product{}
	fmt.Println("Final price:", calc.CalculatePrice("Cheese", 55.67, 2.12))

	x := Product{}
	fmt.Println("x: ", x.CalculatePrice("Butter", 21.11, 1.11))
}
