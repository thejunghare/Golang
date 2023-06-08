package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.youtube.com/"

func main() {
	fmt.Println("Request response")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is type of :%T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
