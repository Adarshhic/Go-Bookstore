# Go Bookstore REST API

A RESTful Bookstore Management API built with **Go (Golang)**, **Gorilla Mux**, and **GORM** connected to a **MySQL** database.

---

## Features

- **Create Book**: Add a new book with name, author, and publication details.
- **Get All Books**: Fetch a list of all books in the database.
- **Get Book by ID**: Retrieve details of a specific book by its ID.
- **Update Book**: Update book details by ID.
- **Delete Book**: Remove a book from the database by ID.
- **Auto Database Migration**: Automatically migrates the database schema on startup using GORM.
- **Configurable Environment**: Supports environment variables for secure database configuration.

---

## Tech Stack

- **Language:** Go (1.20+)
- **Router:** [Gorilla Mux](https://github.com/gorilla/mux)
- **ORM:** [GORM (v1)](https://github.com/jinzhu/gorm)
- **Database:** MySQL
- **Dialect:** `github.com/jinzhu/gorm/dialects/mysql`

---

## Project Structure

```text
Project_2/
├── cmd/
│   └── main/
│       └── main.go              # Application entry point
├── pkg/
│   ├── config/
│   │   └── app.go               # MySQL database connection & GORM setup
│   ├── controllers/
│   │   └── book-controller.go   # HTTP handlers for CRUD operations
│   ├── models/
│   │   └── book.go              # Book model and database queries
│   ├── routes/
│   │   └── bookstore-routes.go  # Route registration
│   └── utils/
│       └── utils.go             # JSON body parser helper
├── .env.example                 # Example environment variables
├── .gitignore                   # Ignored files (binaries, env, etc.)
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies checksums
└── README.md                    # Project documentation
```

---

## Prerequisites

- [Go](https://go.dev/dl/) (version 1.20 or newer)
- [MySQL Server](https://dev.mysql.com/downloads/mysql/) running locally or remotely

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Adarshhic/Go-Bookstore.git
cd Go-Bookstore
```

### 2. Create the MySQL Database

Open your MySQL shell or client and create the database:

```sql
CREATE DATABASE simplerest;
```

### 3. Configure Database Credentials

You can configure the database connection using environment variables or a full DSN connection string:

#### Option A: Set Environment Variables (Recommended)

**PowerShell (Windows):**
```powershell
$env:MYSQL_USER="root"
$env:MYSQL_PASSWORD="your_password"
$env:MYSQL_HOST="127.0.0.1"
$env:MYSQL_PORT="3306"
$env:MYSQL_DBNAME="simplerest"
```

**Bash / Linux / macOS:**
```bash
export MYSQL_USER="root"
export MYSQL_PASSWORD="your_password"
export MYSQL_HOST="127.0.0.1"
export MYSQL_PORT="3306"
export MYSQL_DBNAME="simplerest"
```

#### Option B: Set Full Connection String (`MYSQL_DSN`)

```powershell
$env:MYSQL_DSN="root:your_password@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
```

---

## Running the Application

Download the dependencies and run the server:

```bash
go mod tidy
go run ./cmd/main/main.go
```

The server will start listening on **`http://localhost:9010`**.

---

## API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/book/` | Get all books |
| `GET` | `/book/{bookId}` | Get book by ID |
| `POST` | `/book/` | Create a new book |
| `PUT` | `/book/{bookId}` | Update a book by ID |
| `DELETE` | `/book/{bookId}` | Delete a book by ID |

---

## API Usage Examples

### 1. Create a Book (`POST /book/`)

**Request:**
```bash
curl -X POST http://localhost:9010/book/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "publication": "Addison-Wesley"
  }'
```

**Response (`200 OK`):**
```json
{
  "ID": 1,
  "CreatedAt": "2026-10-09T17:58:05Z",
  "UpdatedAt": "2026-10-09T17:58:05Z",
  "DeletedAt": null,
  "name": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "publication": "Addison-Wesley"
}
```

---

### 2. Get All Books (`GET /book/`)

**Request:**
```bash
curl http://localhost:9010/book/
```

**Response (`200 OK`):**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2026-10-09T17:58:05Z",
    "UpdatedAt": "2026-10-09T17:58:05Z",
    "DeletedAt": null,
    "name": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "publication": "Addison-Wesley"
  }
]
```

---

### 3. Get Book by ID (`GET /book/{bookId}`)

**Request:**
```bash
curl http://localhost:9010/book/1
```

---

### 4. Update a Book (`PUT /book/{bookId}`)

**Request:**
```bash
curl -X PUT http://localhost:9010/book/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Go Programming Language (2nd Edition)",
    "author": "Alan Donovan & Brian Kernighan",
    "publication": "Addison-Wesley"
  }'
```

---

### 5. Delete a Book (`DELETE /book/{bookId}`)

**Request:**
```bash
curl -X DELETE http://localhost:9010/book/1
```

---

## License

This project is open source and available under the [MIT License](LICENSE).
