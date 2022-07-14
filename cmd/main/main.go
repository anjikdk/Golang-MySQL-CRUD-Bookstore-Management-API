package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/anjikdk/go-bookstore/pkg/routes"
	"fmt"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started with port: 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}