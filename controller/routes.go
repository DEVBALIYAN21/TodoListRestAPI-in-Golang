package controller

import "net/http"


func Register() *http.ServeMux{
	mux := http.NewServeMux() //Initialising the mux 
	mux.HandleFunc("/",Data()) //getting direct by hitting the Localhost://portnumber/
	return mux
}