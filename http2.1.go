package main

import (
	"log"
	"fmt"
     "sync"
	"net/http"
	"net/http/httptest"
	"strconv"
)

var (
	currentFib   int = 0
	nextFib      int = 1
	mu           sync.Mutex
	requestCount int
    )
    
func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, strconv.Itoa(currentFib))
	currentFib, nextFib = nextFib, currentFib+nextFib
	requestCount++
	mu.Unlock()
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
    return
}


func main() {
	requestCount = 0

	tt := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{"First Fibonacci Number", http.StatusOK, "0"},
		{"Second Fibonacci Number", http.StatusOK, "1"},
		{"Third Fibonacci Number", http.StatusOK, "1"},
		{"Fourth Fibonacci Number", http.StatusOK, "2"},
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(FibonacciHandler))
	mux.Handle("/metrics", http.HandlerFunc(MetricsHandler))
	
	http.ListenAndServe("127:0.0.1:8080", mux)

	for _, tc := range tt {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				log.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(FibonacciHandler)

			handler.ServeHTTP(rr, req)

			if tc.expectedStatus != rr.Code {
				fmt.Println("failed error code")
			}

			if tc.expectedBody != rr.Body.String() {
				fmt.Println("failed body", rr.Body.String)
			}
	}
}
