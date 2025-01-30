package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type Name struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    log.SetOutput(w)
	n := r.URL.Query().Get("name")
	name := Name{Name: n}
	nj, err := json.Marshal(name)
	if err != nil {
		http.Error(w, "err", http.StatusInternalServerError)
	}
	w.Write(nj)
    log.Println(nj)
    //(next http.HandlerFunc)
	//next(w, r)

}

func main() {
	mux := http.NewServeMux()
    mux.HandleFunc("/hello", HelloHandler)
    http.ListenAndServe(":8080", nil)
}