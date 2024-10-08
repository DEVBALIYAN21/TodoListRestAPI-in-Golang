package controller

import (
	"API2/models"
	"API2/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func Data() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost{										// Method to Insert the Data in the  Database means if post method hitted 
			data:=views.Data{} 												// creating A view
			json.NewDecoder(r.Body).Decode(&data) 							// Decoding the Data before Putting in the Database
			existtodo,err:=models.CheckExist(data.Name,data.Todo)			 //Checking if the todo already exist under same name 
			if err!=nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			if existtodo{ 													// if exists then returning a meassage that todo already exists
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte(fmt.Sprintf("Todo Already Exist "))) 		// simple message 
				return
			}
			if err:=models.Insertdata(data.Name,data.Todo);err!=nil{ 	// calling InsertData function from models 
				w.Write([]byte("Error Occurred "))
			}
			w.WriteHeader(http.StatusOK) 								// Set the status code to 200 (OK)
			w.Write([]byte(fmt.Sprintln("Added Successfully The Data of",data.Name))) // cutom message in Response tab
		} else if r.Method == http.MethodGet{ 							// if  hit get method 
			name := r.URL.Query().Get("name") 							// Checking if Name is Given As parameter 
			if name !=""{ 												// name is not empty 
				exist,err:=models.AlreadyExist(name)					// Checking for Data Exist Under That name 
				if !exist{												// if not then return NO Data Found 
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("No Data Found"))
					return
				}
				data,err:=models.SearchByName(name)						//Searching By Name
			if err!=nil{
				w.Write([]byte("Some Error Occured "))
			}
			json.NewEncoder(w).Encode(data) 							// Encoding the Data In Json Format 
			}else { 													// if name is empty then fetch all the data 
				data,err:= models.GetData() 							// Calling GetData Function to fetch all the data 
				if err!=nil{
					w.Write([]byte ("error occured"))
				}
				json.NewEncoder(w).Encode(data)  						// encoding the data in JSON format 
			}

		} else if r.Method == http.MethodDelete{						 //when delete request hit on the url 
			name:=r.URL.Path[1:]										//Sclicing the delete request extra colun (:) 
			exists,erro:=models.AlreadyExist(name) 						//chekking if the todo exist under that name 
			if erro!=nil{
				http.Error(w, erro.Error(), http.StatusInternalServerError) //if error came return that error 
				}
			if !exists{ 													// if todo is not there then return a message 
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Todo Not Found")) 							// returning the message that todo not exist
				return
			}
			err:=models.DeleteByName(name) 									// calling the function to delete the data by name from models package 
			if err!=nil{
				w.Write([]byte("Some Error Occured "))
				}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf( "Deleted Successfully The Data of  %s",name)))

		}
	}
}