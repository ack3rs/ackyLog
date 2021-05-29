package main

import (
	"fmt"
	"log"
	"net/http"

	myCustomLogger "github.com/acky666/ackyLog"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timed := myCustomLogger.TIMED("New Request")
		defer timed()
		myCustomLogger.WEB("Logging a new Web Request", r)
		myCustomLogger.WEBFORM("Display Anything Sent to this Page", r)

		fmt.Fprintf(w, "Hello!")
	})

	myCustomLogger.INFO("Starting Server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
