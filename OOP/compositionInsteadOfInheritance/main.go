package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

type blogPost struct {
	title   string
	content string
	author // Composition Instead of Inheritance 
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.author.fullName())
	fmt.Println("Bio: ", b.bio)
}

func main() {
	author1 := author{
		"John",
		"Cena",
		"WWE Superstar",
	}

	blogPost1 := blogPost{
		"US champ",
		"How did I become us champ",
		author1,
	}

	blogPost1.details()
}
