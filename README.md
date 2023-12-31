# Basic Book REST API
&rarr; This is a simple RESTful API for managing a collection of books. It is built using Go and the Chi router, with PostgreSQL as the database.

## Prerequisites
&rarr; Before running this application, make sure you have the following installed:

+ Go (https://golang.org/)
  
+ PostgreSQL (https://www.postgresql.org/)
  
+ Install Postgresql driver:
  
  ```bash
    go get github.com/lib/pq
  ```
   
   or

  ```bash
  go get github.com/jmoiron/sqlx
  ```
  
  
+ Install go-chi package:
  
   ```bash
      go get -u github.com/go-chi/chi/v5
  ```
  
    
## Database Configuration
&rarr; The application connects to a PostgreSQL database. Make sure to configure the database connection in the init function of the main package. Update the connection string in the following line:
go

```bash
    db, err = sqlx.Open("postgres", "user=postgres password=Your_postgres_password host=Your_local_host dbname=Your_databse_name sslmode=disable")
```

## API Endpoints
 + GET /getallbooks: Get a list of all books.
 + GET /getbooksbyid/{id}: Get a book by ID.
 + POST /createbooks: Create a new book.
 + DELETE /deletebooksbyid/{id}: Delete a book by ID.
 + PUT /updatebooksbyid/{id}: Update a book by ID.
    
## Usage : 
&rarr; use Postman to interact with the Book REST API
### => Creating a Book
Send a POST request to http://localhost:2222/createbooks with the following JSON payload:

#### json
{
  "id": "1",
  "title": "Sample Book",
  "author": "John Doe",
  "publishedate": "2023-01-01",
  "language": "English"
}

### => Getting All Books
Send a GET request to http://localhost:2222/getallbooks.

### => Getting a Book by ID
Send a GET request to http://localhost:2222/getbooksbyid/{id}, replacing {id} with the actual ID.

### => Updating a Book by ID
Send a PUT request to http://localhost:2222/updatebooksbyid/{id}, replacing {id} with the actual ID. Include the updated book details in the request body.

### => Deleting a Book by ID
Send a DELETE request to http://localhost:2222/deletebooksbyid/{id}, replacing {id} with the actual ID.

## Note:
&rarr; This is a simple example, and error handling, security, and other production considerations are not fully implemented. Ensure proper security measures are in place before deploying this application in a production environment.
