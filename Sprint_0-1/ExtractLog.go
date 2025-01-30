package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {

	var resultstring []string

	f, ferror := os.Open(inputFileName)
	if ferror != nil {
		return nil, ferror
	}
	defer f.Close()

	payload := make(map[time.Time]string)

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		payloadstr, payloaderr := time.Parse("02.01.2006", fileScanner.Text()[:9])
		if payloaderr != nil {
			return nil, payloaderr
		}
		payload[payloadstr] = fileScanner.Text()
	}

	for start.Compare(end) <= 0 {
		if  payload[start] != "" {
		    resultstring = append(resultstring, payload[start])		
		}
		  start = start.AddDate(0,0,1)
	  }

	if resultstring == nil {
		return resultstring, fmt.Errorf("nothing to return")
	}

	return resultstring, nil
}
