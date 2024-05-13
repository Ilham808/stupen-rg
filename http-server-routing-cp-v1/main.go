package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		fmt.Fprintf(w, "%s", currentTime.Format("Monday, 02 January 2006"))
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			fmt.Fprintf(w, "Hello, %s!", name)
		} else {
			fmt.Fprintf(w, "Hello there")
		}
	} // TODO: replace this
}

func main() {
	// TODO: answer here
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())
	http.ListenAndServe("localhost:8080", nil)
}
