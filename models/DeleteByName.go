package models


func DeleteByName(name string) error{
	Query,err:=con.Query("DELETE FROM TODOLIST WHERE NAME = ?",name) //Querry To delete the todo by name
	if err!=nil{
		return err
	}
	defer Query.Close() // closing the query
	return nil
}