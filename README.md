# students-api

A simple API to manage students.

## API Endpoints

| Method | Endpoint                | Description                          |
|--------|-------------------------|--------------------------------------|
| GET    | `/students`             | Retrieve a list of all students.    |
| GET    | `/students/{id}`        | Retrieve a student by ID.            |
| POST   | `/students`             | Create a new student.                |
| PUT    | `/students/{id}`        | Update an existing student by ID.    |
| DELETE | `/students/{id}`        | Delete a student by ID.              |

## Installation and Setup Instructions

> [!NOTE]  
> This project uses sqlite3 as the database

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/geekyharsh05/students-api.git
   cd students-api
   ```

2. **Install Dependencies:**
   Make sure you have Go installed on your machine. You can install the necessary dependencies using:
   ```bash
   go mod tidy
   ```

3. **Run the Application:**
   You can run the application using:
   ```bash
    go run cmd/students-api/main.go -config config/local.yaml
   ```

4. **Access the API:**
   The API will be available at `http://localhost:8082`. You can use tools like Postman or curl to test the endpoints.

## Example Usage

- **Get All Students:**
  ```bash
  curl -X GET http://localhost:8082/students
  ```

- **Get a Student by ID:**
  ```bash
  curl -X GET http://localhost:8082/students/1
  ```

- **Create a New Student:**
  ```bash
  curl -X POST http://localhost:8082/students -d '{"name": "John Doe", "email": "john@example.com", "age": 20}'
  ```

- **Update a Student:**
  ```bash
  curl -X PUT http://localhost:8082/students/1 -d '{"name": "John Smith", "email": "johnsmith@example.com", "age": 21}'
  ```

- **Delete a Student:**
  ```bash
  curl -X DELETE http://localhost:8082/students/1
  ```
