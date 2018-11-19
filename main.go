package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//func handlerFunc(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	//fmt.Fprint(w,r.URL.Path)
//	if r.URL.Path == "/" {
//
//		fmt.Fprint(w, "<h1>Welcome to my site - change</h1>")
//	} else if r.URL.Path == "/contact" {
//
//		fmt.Fprint(w, "To get in touch, please send a email to <a href=\"mailto:support@gmail.com\">support@gmail.com</a>.")
//	} else {
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprint(w, "<h1>Page Not Found </h1>")
//	}
//
//}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site - change</h1>")

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send a email to <a href=\"mailto:support@gmail.com\">support@gmail.com</a>.")
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page Not Found </h1>")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", router)
}
