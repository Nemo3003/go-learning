package main

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type messageToSend struct {
	message 	string
	sender 		user
	recipient 	user
}

type user struct {
	name 	string
	number 	int
}


func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func mainStructs() {
	// Sample data
	sender := user{name: "John", number: 123456789}
	recipient := user{name: "Alice", number: 987654321}

	// Valid message
	validMessage := messageToSend{
		message:   "Hello, Alice!",
		sender:    sender,
		recipient: recipient,
	}

	// Invalid message with empty sender name
	invalidSenderName := messageToSend{
		message:   "Invalid message",
		sender:    user{name: "", number: 123456789},
		recipient: recipient,
	}

	// Invalid message with empty recipient name
	invalidRecipientName := messageToSend{
		message:   "Invalid message",
		sender:    sender,
		recipient: user{name: "", number: 987654321},
	}

	// Invalid message with zero sender number
	invalidSenderNumber := messageToSend{
		message:   "Invalid message",
		sender:    user{name: "John", number: 0},
		recipient: recipient,
	}

	// Invalid message with zero recipient number
	invalidRecipientNumber := messageToSend{
		message:   "Invalid message",
		sender:    sender,
		recipient: user{name: "Alice", number: 0},
	}

	// Testing the function
	fmt.Println("Valid Message:", canSendMessage(validMessage))                // Should print: true
	fmt.Println("Invalid Sender Name:", canSendMessage(invalidSenderName))     // Should print: false
	fmt.Println("Invalid Recipient Name:", canSendMessage(invalidRecipientName)) // Should print: false
	fmt.Println("Invalid Sender Number:", canSendMessage(invalidSenderNumber))   // Should print: false
	fmt.Println("Invalid Recipient Number:", canSendMessage(invalidRecipientNumber)) // Should print: false
}