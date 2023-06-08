package main

import "fmt"

func main() {
	/*
		Maps are used to store data values in key:value pairs.
		syntax -> map_name := map[keyType]valueType{Key1:value1, Key2:value2,...}
	*/

	bike := map[string]string{"Brand": "Honda", "Model": "cb shine"}
	fmt.Println(bike)

	/*
		Create a map using key function
		map_name = make(map[keyType]valueType)
	*/

	my_map := make(map[string]string) // my_map is empty ->  make function is the correct way to create an empty map
	my_map["brand"] = "Honda"         // my_map is no longer empty

	/* 
		Create empty map without using make function
		
		var map_name map[keyType]ValueType

		Note: Create an empty map  fucntion without make function will cause a runtime panic
			panic -> something went unexpectedly wrong
	*/

}
