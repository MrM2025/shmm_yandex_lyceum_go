package tasks

import (
	"fmt"
	"net/http"
	"os/user"
)

type EnsureAuth struct {
	handler http.Handler
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Header.Get("Authorization")
		email := r.URL.Query.Get("email")
		pw := r.URL.Query.Get("password")

		
		if name == "" || email == "" || pw == "" {
		http.Error(w, "please sign-in", http.StatusUnauthorized)
		return
	}

	next(w, r)
	})

}
func answerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The answer is 42")
}

func NewEnsureAuth(handlerToWrap http.Handler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/answer", http.HandlerFunc(answerHandler))	
	wrappedMux := NewEnsureAuth(mux)

	http.ListenAndServe("127:0.0.1:8080", wrappedMux.handler)
}