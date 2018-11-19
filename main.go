package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site - change</h1>")

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send a email to <a href=\"mailto:support@gmail.com\">support@gmail.com</a>.")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", router)
}
