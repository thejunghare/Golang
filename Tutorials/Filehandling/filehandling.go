package main

import (
	"fmt"
	"os"
)

func readFiles() {
	contents, err := os.ReadFile("anotherTest.txt")
	checkError(err)
	fmt.Println(string(contents))
}

func createAndWriteFile() {
	fileName, err := os.Create("say-hello.txt")
	checkError(err)

	content, err := fileName.WriteString("Hello World")
	checkError(err)
	fileName.Close()

	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	readFiles()
	createAndWriteFile()
}
