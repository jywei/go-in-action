// Sample program to show what happens when the outer and inner
// types implement the same interface
package main

// notifier is an interface that defined notification
// type behavior
type notifier interface {
	notify()
}

// user defines an user in the program
type user struct {
	name string
	email string
}

// notify implements a method that can be called via
// a value of type user
func (u *user) notify() {
	fmt.Printf("Send user email to %s<%s>",
		u.name,
		u.email
	)
}

// admin represents an admin user with privilleges
type admin struct {
	user
	level string
}

// notify implements a method that can be called via
// a value of type admin
func (a *admin) notify() {
	fmt.Printf("Send user email to %s<%s>",
		a.name,
		a.email
	)
}

// main is the entry point
func main() {
	// Create an admin user
	ad := admin {
		user: user{
			name: "Roy Wei",
			email: "roywei@email.com"
		},
		level: "super",
	}

	// Send the admin user a notification
	// The embedded inner type's implementation of the
	// interface is NOT "promoted" to the outer type
	sendNotification(&ad)

	// We can access the inner type's method directly
	ad.user.notify()

	// The inner type's method is NOT promoted
	ad.notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications
func sendNotification(n notifier)  {
	n.notify()
}
