/*
	String -> Srquence of characters.
	String -> Implemented differently compared to other languages.
	String in Go -> String in go are slice of bytes.

	Points covered
	1. Length
	2. comparsion
	3. concatenation
	4. Concatenation using %s formater
	7. "RUNE"
	6. Accessing bytes (use %x format specifier)
	7. Accessing bytes using "RUNE"
	8. Accessing Characters (use %c format specifier)
	9. Creating string from slice
	10. Creating string from slice using "RUNE"

*/

package main

import "fmt"

func main() {
	greetings := "Hello"
	str := "Señor"
	fmt.Println(greetings)
	fmt.Println(str)

	fmt.Printf("String: %s\n", greetings) // %s -> format specifier used to print string.

	findsLength(str)

	compareString(str, greetings)

	concatenation(greetings, str)

	concatenationUsingformater(greetings, str)

	accessBytes(greetings)

	fmt.Println("")

	accessCharacters(str) // using %c
	fmt.Println("")

	accessCharactersUsingRune(str) // using runes
	fmt.Println("")

	createStrFromSlice()
	fmt.Println("")

	createStrFromSliceUsingRune()

	illustrateString(greetings)

}

/* Find length of a string */

func findsLength(str string) {
	fmt.Println("Length of string:", len(str))
}

/* String comparsion */

func compareString(str1 string, str2 string){
	if str1 == str2 {
		fmt.Println("String are equal")
	} else {
		fmt.Println("String are not equal")
	}
}

/* String concatenation using "+" */

func concatenation(str1 string, str2 string){
	fmt.Println(str1 + " " + str2)
}

/* String concatenation using %s */
func concatenationUsingformater(str1 string, str2 string){
	fmt.Printf("%s %s", str1 , str2 + "\n")
}

/* Accessing individual bytes of string */

func accessBytes(str string) {
	fmt.Print("Printing individual bytes: ")

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i]) // %x -> format specifier used to print hexadecimal.
	}
}

/* Accessing individual character */

func accessCharacters(str string) {
	fmt.Print("Accessing individual characters: ")

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i]) // %c -> format specifier used to print the characters of the string.
	}
}

/* Accessing characters using "RUNE" */

func accessCharactersUsingRune(str string) {
	fmt.Print("Accessing characters using RUNE: ")

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func illustrateString(str string) {
	name := "Jack"
	age := 25

	fmt.Printf("\nName is %s and age is %d", name, age)
}

/* Creating string from a slice */

func createStrFromSlice() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Print(str)
}

/* Creating string from a slice from RUNE */

func createStrFromSliceUsingRune() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Print(str)
}
