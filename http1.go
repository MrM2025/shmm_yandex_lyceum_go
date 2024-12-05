package main

import (
	"fmt"
	"net/http"
)

func IsLatin(input string) bool {
	for i := range len(input) {
		inpt := []rune(input)[i]
		if inpt > 122 || inpt < 65 || inpt == 91 || inpt == 92 || inpt == 93 || inpt == 94 || inpt == 95 || inpt == 96 {
			return false
		}
	}
	return true
}

func StrangerHandler(w http.ResponseWriter, r *http.Request)  {
	name := r.URL.Query().Get("name")

	if len(name) == 0{
		fmt.Fprintf(w, "hello stranger")
		return
	} 
	if IsLatin(name) == false {
			fmt.Fprintf(w, "hello dirty hacker")
			return
	} else {
		fmt.Fprintf(w, "hello %s", name)
		return
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
    
	http.ListenAndServe(":8080", nil)
    }


