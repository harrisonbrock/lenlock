package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//fmt.Fprint(w,r.URL.Path)
	if r.URL.Path == "/" {

		fmt.Fprint(w, "<h1>Welcome to my site - change</h1>")
	} else if r.URL.Path == "/contact" {

		fmt.Fprint(w, "To get in touch, please send a email to <a href=\"mailto:support@gmail.com\">support@gmail.com</a>.")
	} else {
		fmt.Fprint(w, "<h1>404 Page Not Found </h1>")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
