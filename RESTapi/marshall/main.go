package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Professor struct {
	Name       string     `json:"name"`
	ScienseID  int        `json:science_id`
	IsWorking  bool       `json:is_working`
	University University `json:university`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor{
		Name:      "Karl",
		ScienseID: 245,
		IsWorking: true,
		University: University{
			Name: "MIT",
			City: "Boston",
		},
	}

	byteArr, err := json.MarshalIndent(prof1, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))
}
