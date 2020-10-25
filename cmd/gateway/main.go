package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/greeting", func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(200)
		response.Header().Add("Content-Type", "application/json")
		body, _ := json.Marshal("Hi, John!!!")
		_, _ = response.Write(body)

	})

	log.Fatal(http.ListenAndServe(":9000", nil))

}
