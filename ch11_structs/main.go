package main

import (
	"fmt"
	"strings"
)

// Structs {Structures}
// Keywords: type and struct
type Book struct {
	Title   string
	Author  string
	Subject string
	BookID  int
	Read    bool
}

func main() {
	book1 := Book{
		Title:   "Range",
		Author:  "King",
		Subject: "novel",
		BookID:  0123,
		Read:    true,
	}

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Title: ", book1.Title)
	fmt.Println("Author: ", book1.Author)
	fmt.Println("Title: ", book1.Subject)
	fmt.Println("BookID: ", book1.BookID)
	fmt.Println("Read: ", book1.Read)
	fmt.Println(strings.Repeat("=", 10))
}
