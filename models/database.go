package models

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func ConnectToDB() *sql.DB{
	db,err:=sql.Open("mysql","root:Dev@123@tcp(localhost:3306)/TODO")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	con=db
	return db
}