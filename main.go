package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":3000", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my API")
}