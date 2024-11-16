package main

import (
	"encoding/json"
)

type main struct {
	Student
}

type Student struct { 
	Name string `json:"name"`
	Grade int `json:"grade"`
} 

func modifyJSON(jsonData []byte) ([]byte, error) {
	var studentcard Student

	err := json.Unmarshal(jsonData, &studentcard)
	if err != nil {
		return []byte{}, err
	}

	studentcard.Grade += 2

	cardinbytes, writerr := json.Marshal(studentcard)
	if writerr != nil {	
		return []byte{}, writerr	
	}

	return cardinbytes, nil
}