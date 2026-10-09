package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adarshhic/Go-Bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on localhost:9010...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
