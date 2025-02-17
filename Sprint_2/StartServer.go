package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func ReadSourceHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/provideData")
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  log.Printf("error with reading from resourse: %v", err)
	  return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
	
}

func StartServer(maxTimeout time.Duration) {
	h := http.TimeoutHandler(http.HandlerFunc(ReadSourceHandler), maxTimeout, "max timeout")
	http.Handle("/readSource", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}

}