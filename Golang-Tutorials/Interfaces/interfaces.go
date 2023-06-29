package main

import "fmt"

type Animal interface {
	Speak() string
	Walk() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func (d Dog) Walk() string {
	return "walk"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func (c Cat) Walk() string {
	return "Meow"
}

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}