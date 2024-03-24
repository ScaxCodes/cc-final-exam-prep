package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	Name		string
	Email		string
	Age			int
	Groups	[]string // list of groups the contact belongs to
}

func sortContacts(contacts []Contact) []Contact {
	
	// Custom sort function
	sortContacts := func(i, j int) bool {
		// Sort by the number of groups (descending)
		if len(contacts[i].Groups) != len(contacts[j].Groups) {
			return len(contacts[i].Groups) > len(contacts[j].Groups)
		}
		// If the number of groups is the same, sort by age (ascending)
		if contacts[i].Age != contacts[j].Age {
			return contacts[i].Age < contacts[j].Age
		}
		// If the age is the same, sort alphabetically by name (ascending)
		return contacts[i].Name < contacts[j].Name
	}

	outputSlice := contacts
	// Sort the contacts using the custom sort function
	sort.Slice(outputSlice, sortContacts)
	return outputSlice
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

	contacts = []Contact{
    {
        Name:   "John Doe",
        Email:  "johndoe@example.com",
        Age:    35,
        Groups: []string{"Friends", "Work", "Gym", "Book Club"},
    },
    {
        Name:   "Alice Smith",
        Email:  "alicesmith@example.com",
        Age:    30,
        Groups: []string{"Work", "Gym"},
    },
    {
        Name:   "Charlie Brown",
        Email:  "charliebrown@example.com",
        Age:    28,
        Groups: []string{"Book Club", "Chess Club"},
    },
    {
        Name:   "Bob White",
        Email:  "bobwhite@example.com",
        Age:    45,
        Groups: []string{"Friends", "Chess Club"},
    },
    {
        Name:   "Jane Doe",
        Email:  "janedoe@example.com",
        Age:    32,
        Groups: []string{"Work"},
    },
    {
        Name:   "Mary Johnson",
        Email:  "maryjohnson@example.com",
        Age:    35,
        Groups: []string{"Friends", "Work", "Gym", "Book Club"},
    },
    {
        Name:   "Emma Brown",
        Email:  "emmabrown@example.com",
        Age:    25,
        Groups: []string{},
    },
	}
	result = sortContacts(contacts)
	fmt.Println(result)

	contacts = []Contact{
    {
        Name:   "John Doe",
        Email:  "johndoe@example.com",
        Age:    25,
        Groups: []string{"Friends", "Work", "Gym"},
    },
    {
        Name:   "Jane Smith",
        Email:  "janesmith@example.com",
        Age:    25,
        Groups: []string{"Friends", "Gym"},
    },
    {
        Name:   "Alice Johnson",
        Email:  "alicejohnson@example.com",
        Age:    30,
        Groups: []string{"Friends", "Work"},
    },
	}
	result = sortContacts(contacts)
	fmt.Println(result)

	contacts = []Contact{
    {
        Name:   "Alice",
        Email:  "alice@example.com",
        Age:    20,
        Groups: []string{},
    },
    {
        Name:   "Bob",
        Email:  "bob@example.com",
        Age:    25,
        Groups: []string{"Friends"},
    },
	}
	result = sortContacts(contacts)
	fmt.Println(result)
	
	contacts = []Contact{}
	result = sortContacts(contacts)
	fmt.Println(result)
}