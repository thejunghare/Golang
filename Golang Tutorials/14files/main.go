package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	/* Content for file */
	content := "This text will be added to file"

	/* Creating file */
	file, err := os.Create("./file.txt")
	checkError(err)

	/* Find length of char inside the file */
	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	createFile()

	readFile("./file.txt")
}

/* Read file */
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkError(err)

	fmt.Println("Text data inside the file \n", string(databyte))
}

func createFile() {
	_, err := os.Create("./paddy.txt")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
