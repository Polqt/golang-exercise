package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Polqt/book-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}