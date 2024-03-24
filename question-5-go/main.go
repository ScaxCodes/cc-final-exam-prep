package main

import "fmt"

type Contact struct {
	Name		string
	Email		string
	Age			int
	Groups	[]string // list of groups the contact belongs to
}

func sortContacts(contacts []Contact) []Contact {
	
}

func main() {
	contacts := []Contact{
		Contact{
					 Name:   "John Doe",
					 Email:  "johndoe@example.com",
					 Age:    30,
					 Groups: []string{"Friends", "Work"},
		}, Contact{
					 Name:   "Jane Doe",
					 Email:  "janedoe@example.com",
					 Age:    25,
					 Groups: []string{"Work"},
		}, Contact{
					 Name:   "Alice Johnson",
					 Email:  "alicejohnson@example.com",
					 Age:    30,
					 Groups: []string{"Friends", "Work"},
		},
	}
	result := sortContacts(contacts)
	fmt.Println(result)
}