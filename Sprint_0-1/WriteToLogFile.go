package main

import (
	"log"
	"os"
)

func WriteToLogFile(message string, fileName string) error{
	f, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("hello world")
	return nil

}