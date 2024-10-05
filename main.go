package main

import (
	"API2/controller"
	"API2/models"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux:=controller.Register()
	db:=models.ConnectToDB() //strting the DB connection
	defer db.Close() // closing the database at the end of the server
	port:=1999
	fmt.Println("Running at ",port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux) // starting at port 1999
}