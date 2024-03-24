package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Sales int 
}

func topAuthors(books []Book) map[string]int {
	authorSales := make(map[string]int)

	for _, book := range books {
    _, ok := authorSales[book.Author]
    if !ok {
        authorSales[book.Author] = 0
    }
    authorSales[book.Author] += book.Sales
	}

	for author, sales := range authorSales {
		if sales < 10000 { delete(authorSales, author) }
	}

	return authorSales
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

	books = []Book{
    {
        Title:  "The Great Escape",
        Author: "John Doe",
        Sales:  4500,
    },
    {
        Title:  "The Lost World",
        Author: "Michael Crichton",
        Sales:  20000,
    },
    {
        Title:  "The Final Countdown",
        Author: "John Doe",
        Sales:  8000,
    },
    {
        Title:  "The Tale of Two Cities",
        Author: "Charles Dickens",
        Sales:  5000,
    },
    {
        Title:  "Under the Dome",
        Author: "Stephen King",
        Sales:  25000,
    },
    {
        Title:  "Pet Sematary",
        Author: "Stephen King",
        Sales:  8000,
    },
	}
	result = topAuthors(books)
	fmt.Println(result)

	books = []Book{
    {
        Title:  "Book 1",
        Author: "Author 1",
        Sales:  8000,
    },
    {
        Title:  "Book 2",
        Author: "Author 2",
        Sales:  6000,
    },
    {
        Title:  "Book 3",
        Author: "Author 3",
        Sales:  12000,
    },
    {
        Title:  "Book 4",
        Author: "Author 1",
        Sales:  5000,
    },
    {
        Title:  "Book 5",
        Author: "Author 3",
        Sales:  1000,
    },
    {
        Title:  "Book 6",
        Author: "Author 4",
        Sales:  15000,
    },
	}
	result = topAuthors(books)
	fmt.Println(result)

	books = []Book{
    {
        Title:  "Book 1",
        Author: "Author 1",
        Sales:  3000,
    },
    {
        Title:  "Book 2",
        Author: "Author 3",
        Sales:  4000,
    },
    {
        Title:  "Book 3",
        Author: "Author 1",
        Sales:  2000,
    },
    {
        Title:  "Book 4",
        Author: "Author 3",
        Sales:  7000,
    },
	}
	result = topAuthors(books)
	fmt.Println(result)

	books = []Book{
    {
        Title:  "Book 1",
        Author: "Author 1",
        Sales:  5000,
    },
    {
        Title:  "Book 2",
        Author: "Author 2",
        Sales:  2000,
    },
	}	
	result = topAuthors(books)
	fmt.Println(result)

	books = []Book{}
	result = topAuthors(books)
	fmt.Println(result)
}