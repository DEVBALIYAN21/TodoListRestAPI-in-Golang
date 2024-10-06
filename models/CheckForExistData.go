package  models


func CheckExist(name,todo string )(bool,error){
	// check with database 
	rows,err:=con.Query("SELECT 1 FROM TODOLIST WHERE NAME=? AND TODO =?",name ,todo) // Querry for Seaching the Todo and name 
	if err!=nil{
		return false,err // if error occured then return the Error 
	}
	defer rows.Close() // closing the Query
	if rows.Next(){ //seaching if the data Exists
		return true,nil //if Exists Then Returning true 
	}
	return false,nil // else Returning False 
	
} 


func AlreadyExist(name string) (bool,error){
	rows,err:=con.Query("SELECT 1 FROM TODOLIST WHERE NAME=?",name)//checking if a todo exist for that particular name 
	if err!=nil{
		return false,err //if error then return Error 
	}
	if rows.Next(){
		return true,nil //if exist then return true
	}
	return false,nil //else return false
}