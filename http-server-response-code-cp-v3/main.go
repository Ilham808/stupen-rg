package main

import (
	"fmt"
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			name := request.URL.Query().Get("name")
			if name == "" {
				writer.WriteHeader(http.StatusNotFound)
				writer.Write([]byte("Data not found"))
				return
			}

			checkName := IsNameExists(name)
			if !checkName {
				writer.WriteHeader(http.StatusNotFound)
				writer.Write([]byte("Data not found"))

				return
			}
			fmt.Fprintf(writer, "Name is exists")
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method is not allowed"))
			return
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
