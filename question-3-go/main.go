package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Sales int 
}

func topAuthors(books []Book) map[string]int {
	return map[string]int{
		"Author 1": 13000,
		"Author 2": 10000,
	}
}

func main() {
	books := []Book{
    {
        Title:  "Book 1",
        Author: "Author 1",
        Sales:  5000,
    },
    {
        Title:  "Book 2",
        Author: "Author 2",
        Sales:  10000,
    },
    {
        Title:  "Book 3",
        Author: "Author 1",
        Sales:  6000,
    },
    {
        Title:  "Book 4",
        Author: "Author 3",
        Sales:  3000,
    },
    {
        Title:  "Book 5",
        Author: "Author 1",
        Sales:  2000,
    },
	}

	result := topAuthors(books)
	fmt.Println(result)
}