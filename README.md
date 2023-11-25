# Book REST API
This is a simple RESTful API for managing a collection of books. It is built using Go and the Chi router, with PostgreSQL as the database.
# Prerequisites
Before running this application, make sure you have the following installed:
    Go (https://golang.org/)
    PostgreSQL (https://www.postgresql.org/)
    
# Database Configuration
The application connects to a PostgreSQL database. Make sure to configure the database connection in the init function of the main package. Update the connection string in the following line:
go

``db, err = sqlx.Open("postgres", "user=postgres password=admin1234 host=127.0.0.1 dbname=book_rest_api sslmode=disable")
``
# API Endpoints
    GET /getallbooks: Get a list of all books.
    GET /getbooksbyid/{id}: Get a book by ID.
    POST /createbooks: Create a new book.
    DELETE /deletebooksbyid/{id}: Delete a book by ID.
    PUT /updatebooksbyid/{id}: Update a book by ID.

