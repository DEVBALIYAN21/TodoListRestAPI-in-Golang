package models


func DeleteByName(name string) error{
	Query,err:=con.Query("DELETE FROM TODOLIST WHERE NAME = ?",name)
	if err!=nil{
		return err
	}
	defer Query.Close()
	return nil
}