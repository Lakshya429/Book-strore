# Golang Book Store

## Overview
The **Golang Book Store** is a web application built using Go (Golang) that allows users to manage a collection of books. It provides RESTful APIs for adding, updating, retrieving, and deleting books from the store.

## Features
- Add new books with title, author, price, and publication year.
- Update book details.
- Retrieve a list of all books.
- Fetch details of a specific book by ID.
- Delete a book from the store.

## Technologies Used
- **Golang** – Backend development.
- **Gin** – Web framework for handling API routes.
- **GORM** – ORM for database interactions.
- **PostgreSQL** – Database for storing book records.
- **Docker** – Containerization for deployment.

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- Golang (v1.20+ recommended)
- PostgreSQL
- Docker (optional, for containerized setup)

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/golang-bookstore.git
   cd golang-bookstore
   ```

2. Create a `.env` file and configure database credentials:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=bookstore
   ```

3. Install dependencies:
   ```sh
   go mod tidy
   ```

4. Run database migrations:
   ```sh
   go run migrations/migrate.go
   ```

5. Start the server:
   ```sh
   go run main.go
   ```

6. The API will be available at `http://localhost:8080`

## API Endpoints
| Method | Endpoint        | Description               |
|--------|----------------|---------------------------|
| GET    | `/books`       | Get all books             |
| GET    | `/books/:id`   | Get book by ID            |
| POST   | `/books`       | Add a new book            |
| PUT    | `/books/:id`   | Update book details       |
| DELETE | `/books/:id`   | Delete a book             |

