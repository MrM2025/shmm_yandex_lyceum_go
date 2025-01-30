package main

import (
	"time"
	"fmt"
    "sync"
	"net/http"
)

func StartServer(t time.Duration) {
	http.ListenAndServe(":8080", nil)
    }


	var (
		currentFib   int
		nextFib      int = 1
		mu           sync.Mutex
		requestCount int
	    )
	    
	    func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintln(w, currentFib)
		currentFib, nextFib = nextFib, currentFib+nextFib
		requestCount++
		mu.Unlock()
	    }


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
    
	http.ListenAndServe(":8080", nil)
    }

