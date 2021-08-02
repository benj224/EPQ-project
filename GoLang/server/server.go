package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

//basic handling function to return contents of test.json
func ReturnJSON(w http.ResponseWriter, r *http.Request) {
	jsonFile, _ := os.Open("test.json")
	byteJSON, _ := ioutil.ReadAll(jsonFile)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteJSON)
}

//start and listen to a http server
func main() {
	http.HandleFunc("/", ReturnJSON)
	http.ListenAndServe(":8080", nil)
}
