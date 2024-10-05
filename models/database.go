package models

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB //Creating this variable just to Make a connection only in this file 

// Just Improving the Security  
// If the con is started with capital C Then Every File can Access That Variable 
// Keep In Mind This Thing 

func ConnectToDB() *sql.DB{  // Returning the Instance of the database 
	db,err:=sql.Open("mysql","root:Dev@123@tcp(localhost:3306)/TODO") // format to connect with the database 
	// sql.Open("Sql Type Example -> mysql ","Username:Password@(tcp:server:port)/Database Name ")
	if err!=nil{
		log.Fatal(err)
	}
	con=db // Creating a secured Connection  Without this we can connect also
	
	fmt.Println("Connected to The Database ....")
	return db // returning the connection or instance of db
}