package main

import (
	"encoding/json"
	"fmt"
)

type Pupil struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {

	var pupils, mergedpupilslist []Pupil
	var mergedJSON []byte

	for _, value := range jsonDataList {

		err := json.Unmarshal(value, &pupils)
		if err != nil {
			return nil, err
		}
		mergedpupilslist = append(mergedpupilslist, pupils...)
	}

	mergedJSON, uerror := json.Marshal(mergedpupilslist)
	if uerror != nil {
		return nil, uerror
	}

	return mergedJSON, nil
}

func main() {
	inputJSON1 := []byte(`[ 
  { 
   "name": "Oleg", 
   "class": "9B" 
  }, 
  { 
   "name": "Ivan", 
   "class": "9A" 
  } 
 ]`)

	inputJSON2 := []byte(`[
	  {
	   "name": "Maria",
	   "class": "9B"
	  },
	  {
	   "name": "John",
	   "class": "9A"
	  }
	 ]`)

	inputJSON3 := []byte(`[    
	
 ]`)

	result, jerror := mergeJSONData(inputJSON1, inputJSON2, inputJSON3)
	if jerror == nil {
		fmt.Println(string(result))
	} else {
		fmt.Println(jerror)
	}
}
