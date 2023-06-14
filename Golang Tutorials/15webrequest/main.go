package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://www.youtube.com/"

func main() {
	fmt.Println("Request response")

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
