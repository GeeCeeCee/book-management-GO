package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GeeCeeCee/bookings/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	

	hostname:= "localhost:"+os.Getenv("PORT")

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(hostname, r))
}





