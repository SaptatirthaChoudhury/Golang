package main

import (
	"fmt"
)

// Define the struct

type User struct {
	Username    string
	Email       string
	SignInCount int
	IsActive    bool
}

type Mohor struct {
	Username string
	Email    string
}

func initUserWithField() {
	user1 := User{
		Username:    "John",
		Email:       "wick@protonmail.com",
		SignInCount: 1,
		IsActive:    true,
	}

	fmt.Println("User1 :", user1)
}

func initWithPointer() {
	user2 := &User{
		Username: "Sam",
		Email:    "altman@gmail.com",
	}

	fmt.Println("User2 :", *user2)
}

// Value Receiver Method
// This does NOT modify the original struct because it works on a copy.
func (u Mohor) PrintInfo() {
	fmt.Printf("Mohor : %s, Email: %s\n", u.Username, u.Email)
}

// Pointer Receiver Method
// This CAN modify the original struct because it works with a pointer.
func (u *Mohor) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	initUserWithField()
	initWithPointer()

	// Create a User instance
	user3 := Mohor{
		Username: "Tom",
		Email:    "tom@yahoo.com",
	}

	fmt.Println("=== Before Update ===")
	user3.PrintInfo()

	// Call pointer receiver method
	user3.UpdateEmail("tom@gmail.com")

	fmt.Println("=== After Update ===")
	user3.PrintInfo()
}
