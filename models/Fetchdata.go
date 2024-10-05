package models

import "API2/views"

func GetData() ([] views.Data, error) {
	rows, err := con.Query("SELECT * FROM TODOLIST")
	if err != nil {
		return nil,err
	}
	todos := []views.Data{}
	for rows.Next() {
		data:=views.Data{}
		rows.Scan(&data.Name,&data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}