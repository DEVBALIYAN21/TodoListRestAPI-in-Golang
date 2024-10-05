

# 📝 Todo API

This is a simple Todo API that allows you to manage a todo list for different users. The API supports creating, retrieving, and deleting todos while ensuring that duplicate todos for the same user are not allowed.

## 📬 Post Request Example
```json
{
  "name": "dev",
  "todo": "Get Milk From the Market"
}
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
