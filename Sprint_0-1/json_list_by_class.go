package main

import (
	"encoding/json"
	"fmt"
)

type Pupil struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var pupils []Pupil

	grouppedpupils := make(map[string][]Pupil)
	listbyclass := make(map[string][]byte)
	uerror := json.Unmarshal(jsonData, &pupils)
	if uerror != nil {
		return nil, uerror
	}

	for _, value := range pupils {
		grouppedpupils[value.Class] = append(grouppedpupils[value.Class], value)
	}

	fmt.Println(grouppedpupils)

	for key, value := range grouppedpupils {

		pupilasbytearray, merror := json.Marshal(value)
		if merror != nil {
			return nil, merror
		}
		fmt.Println(string(pupilasbytearray))
		listbyclass[key] = append(listbyclass[key], pupilasbytearray...)
	}

	return listbyclass, nil
}

func main() {

	inputJSON := []byte(`[
		{
		    "name": "Oleg",
		    "class": "9B"
		},
		{
		    "name": "Ivan",
		    "class": "9A"
		},
		{
		    "name": "Maria",
		    "class": "9B"
		},
		{
		    "name": "John",
		    "class": "9A"
		}
	 ]`)

	result, jerror := splitJSONByClass(inputJSON)
	if jerror == nil {
		for key, value := range result {
			fmt.Println(key, string(value))
		}
	} else {
		fmt.Println(jerror)
	}
}
