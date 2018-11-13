package main

import (
	"fmt"
	"net/http"
)

func demoRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site - change</h1>")
}

func main() {
	http.HandleFunc("/", demoRoute)
	http.ListenAndServe(":3000", nil)
}
