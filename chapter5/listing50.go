// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type
package main

import "fmt"

// user defines an user in the program
type user struct {
	name  string
	email string
}

// notify implements a methd that can be called via
// a value of type user
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>",
		u.name,
		u.email)
}

// admin represents an admin user with privileages
type admin struct {
	user  // Embedded Type
	level string
}

// main is the entry point
func main() {
	// Create an admin user
	ad := admin{
		user: user{
			name:  "Roy Wei",
			email: "roy.wei@email.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly
	ad.user.notify()

	// The inner type's method is promoted
	ad.notify()
}
