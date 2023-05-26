package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/goLinuxCloud", handleReq)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	response["message"] = "Hello GoLinuxCloud Members!"
	response["time"] = "10/26/2022"

	// marshall json
	json, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error JSON marshal: %s", err)
	}
	w.Write(json)
}
