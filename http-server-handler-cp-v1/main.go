package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		weekday := now.Weekday().String()
		day := now.Day()
		month := now.Month().String()
		year := now.Year()
		timeFormatted := fmt.Sprintf("%s, %d %s %d", weekday, day, month, year)
		writer.Write([]byte(timeFormatted))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
