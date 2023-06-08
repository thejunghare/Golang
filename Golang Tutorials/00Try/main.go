package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/* Content to insert in file */
	content := "hello"

	/* create file */
	file, err := os.Create("./file.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println(length)

	defer file.Close()
}

func checkError(err error) {
	panic(err)
}
