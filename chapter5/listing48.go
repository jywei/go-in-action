// Sample program to show how polymorphic behavior with interfaces
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior
type notifier interface {
	notify()
}

// user defines a user in the program
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>",
		u.name,
		u.email)
}

// admin defines an admin in the program
type admin struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>",
		a.name,
		a.email)
}

// main is the entry point
func main() {
	// Create a user value and pass it to sendNotification
	roy := user{"Roy", "roy@email.com"}
	sendNotification(&roy)

	// Create an admin value and pass it to sendNotification
	wei := admin{"wei", "wei@email.com"}
	sendNotification(&wei)
}

// sendNotification accepts values that implement the notifier
// interface and send notifications
func sendNotification(n notifier) {
	n.notify()
}
