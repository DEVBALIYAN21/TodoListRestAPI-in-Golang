package models

import "API2/views"

func GetData() ([] views.Data, error) {
	rows, err := con.Query("SELECT * FROM TODOLIST") // fetching all the details row-wise
	if err != nil {
		return nil,err
	}
	todos := []views.Data{} //creting an instance of the structure to return all the data as an array
	for rows.Next() {
		data:=views.Data{} // now create a single instace for  a single row 
		rows.Scan(&data.Name,&data.Todo) 
		todos = append(todos, data)// putting each row in array of todos
	}
	return todos, nil
}