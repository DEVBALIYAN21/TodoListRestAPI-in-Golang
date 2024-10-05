package models

func Insertdata(name,todo string) error{
	InsertQ,err:=con.Query("INSERT INTO TODOLIST VALUES(?,?)",name,todo)
	defer InsertQ.Close()
	if err!=nil{
		return err
	}
	return nil
}