package main

import (
	"github.com/gorilla/mux"
	"lenslock.com/controllers"
	"net/http"
)



func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUser()

	router := mux.NewRouter()
	router.Handle("/", staticC .Home)
	router.Handle("/contact", staticC.Contact)
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":3000", router)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
