package main

import (
	"fmt"
	"net/http"
)

func mainPaige(res http.ResponseWriter, req *http.Request) {
	headers := req.Header
	headersString := ""
	for key, values := range headers {
		for _, value := range values {
			headersString += fmt.Sprintf("%s: %s\n", key, value)
		}
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(headersString))
	fmt.Println("Incoming request. Client headers:")
	fmt.Println(headersString)
}

func main() {
	fmt.Printf("Starting server.\n")
	http.HandleFunc(`/`, mainPaige)
	err := http.ListenAndServe(`:80`, nil)
	if err != nil {
		panic(err)
	}
}
