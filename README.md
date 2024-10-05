

# 📝 Todo API

This is a simple Todo API that allows you to manage a todo list for different users. The API supports creating, retrieving, and deleting todos .

## 📬 Post Request Example
```json
{
  "name": "dev",
  "todo": "Get Milk From the Market"
}
```
To get started with the project, follow these steps:
Here’s the markdown (MD) code for your SQL database and table setup:


# 🗄️ SQL Setup for Todo API

First Create Tables in MySql

## 1. Create the Database
```sql
CREATE DATABASE TODO;
```

## 2. Switch to the `TODO` Database
```sql
USE TODO;
```

## 3. Create the `TODOLIST` Table
```sql
CREATE TABLE TODOLIST (
    name VARCHAR(100) NOT NULL,
    todo VARCHAR(255) NOT NULL,
);
```

## 🚀 Setting Up the Project

1. **Install the MySQL Driver**:  
   Use the following command to install the MySQL server driver for Go:
   ```bash
   go get -u github.com/go-sql-driver/mysql
   ```

2. **Clean Up Dependencies**:  
   Run the following command to clean up and manage your Go module dependencies:
   ```bash
   go mod tidy
   ```

3. **Run the Application**:  
   Finally, to run the project, use:
   ```bash
   go run main.go
   ```

## 🚀 API Endpoints & Commands(  if Any Error happens then remove .exe from curl)

- **Get All Todos**:  
  ```bash
  curl.exe -X GET http://localhost:1999/
  ```

- **Add a Todo**:  
  ```bash
  curl.exe -X POST http://localhost:1999/ -H "Content-Type: application/json" -d '{"name":"dev", "todo":"Get Milk From the Market"}'
  ```
- **Delete Todos of a User**:  
  ```bash
  curl.exe -X DELETE http://localhost:1999/dev
  ```
- **Fetch Todos for a Specific User**:  
  ```bash
  curl.exe -X GET http://localhost:1999/?name=dev
  ```

## 🧐 Features

Here are some of the project's best features:

- **📝 Get Todos**: Retrieve all todos or filter by a specific user.
- **❌ Delete a Todo**: Remove a todo for a user.
- **🛠️ Tech Stack**:
  - **Programming Language**: Go (Golang)
  - **Database**: SQLite / MySQL
  - **API Framework**: Built using Go's native `net/http` package.

## 🌐 API Endpoints

- **`GET /`**: Retrieves all the todos from the database.
- **`POST /`**: Adds a new todo to the database.
- **`DELETE /{name}`**: Deletes all todos for the specified user.
- **`GET /?name={name}`**: Retrieves todos for the specified user.

## 💻 Built with

Technologies used in the project:

- **Go Lang** 🐹
