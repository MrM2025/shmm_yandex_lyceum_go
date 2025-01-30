package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Answer struct {
	Greetings string `json:"greetings,omitempty"`
	Name      string `json:"name,omitempty"`
}

const Greetings = "hello"

type RPCS struct {
	Status string `json:"status"`
	Result Answer `json:"result,omitempty"`
}

type RPCS1 struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func IsLatin(input string) bool {
	for i := range len(input) {
		inpt := []rune(input)[i]
		if inpt > 122 || inpt < 65 || inpt == 91 || inpt == 92 || inpt == 93 || inpt == 94 || inpt == 95 || inpt == 96 {
			return false
		}
	}
	return true
}

func SetDefaultName(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			res := RPCS{Status: "ok", Result: Answer{Greetings: "hello", Name: "stranger"}}
			ans, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "error - marshaling", http.StatusInternalServerError)
			}
			w.Write(ans)

		}

		h(w, r)
	})
}

func Sanitize(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if !IsLatin(name) {
			panic("")
		}
		h(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if IsLatin(name) && len(name) > 0 {
		res := RPCS{Status: "ok", Result: Answer{Greetings: "hello", Name: name}}
		ans, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "error - marshaling", http.StatusInternalServerError)
		}
		w.Write(ans)

	}
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				res := RPCS{Status: "error", Result: Answer{}}
				ans, err := json.Marshal(res)
				if err != nil {
					http.Error(w, "error - marshaling", http.StatusInternalServerError)
				}
				w.Write(ans)
				return
			}
		}()
		next(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
