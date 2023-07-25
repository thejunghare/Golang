package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Tank struct {
	Name     string   `json:"name"`
	Position Point    `json:"position"`
	Drivers  []string `json:"drivers"`
}

func main() {
	myTank := Tank{
		Name: "tank_1234",
		Position: Point{
			X: 12,
			Y: 34,
		},
		Drivers: []string{"man_1", "man_2"},
	}

	jsondata, err := json.Marshal(&myTank)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(jsondata))

	/* var newTank Tank
	err = json.Unmarshal(jsondata, &newTank)
	fmt.Println(newTank) */
}
