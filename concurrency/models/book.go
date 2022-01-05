package models

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t%q\n"+
			"Author:\t%q\n",
		b.Title, b.Author)
}

func GetBooks() []Book {
	return books
}

var books = []Book{
	{
		ID:     1,
		Author: "Vladimir Nabokov",
		Title:  "Lolita",
	},
	{
		ID:     2,
		Author: "Dan Brown",
		Title:  "Deception Point",
	},
	{
		ID:     3,
		Author: "Charles Dickens",
		Title:  "Great Expectations",
	},
	{
		ID:     4,
		Author: "O. V. Vijayan",
		Title:  "Khasakkinte Ithihasam",
	},
	{
		ID:     5,
		Author: "Paulo Coelho",
		Title:  "The Alchemist",
	},
	{
		ID:     6,
		Author: "Chetan Bhagat",
		Title:  "Five Point Someone",
	},
	{
		ID:     7,
		Author: "Amish Tripathi",
		Title:  "The Immortals of Meluha",
	},
	{
		ID:     8,
		Author: "Gabriel Garcia Marquez",
		Title:  "One hundred years of solitude",
	},
	{
		ID:     9,
		Author: "Robin Sharma",
		Title:  "The Monk who sold his Ferrari",
	},
	{
		ID:     10,
		Author: "Arthur Conan Doyle",
		Title:  "A study in Scarlet",
	},
}
