package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// begin web page
	var hostPlatform = os.Getenv("HOST_PLATFORM")
	var backColor = os.Getenv("BACK_COLOR")

	var htmlHeader = "<!DOCTYPE html><html><h2>Simple web app</h2><h3>Made more complex</h3>test2<br /> TEST TEST TEST"
	fmt.Fprintf(w, htmlHeader)
	fmt.Fprintf(w, "<body style=background-color:%s><p>Running on: %s</p></body></html>", backColor, hostPlatform)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
