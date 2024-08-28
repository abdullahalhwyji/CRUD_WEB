# GoLang Native CRUD Website with MySQL and Bootstrap 5

This repository provides a practical case study for learning native Go (Golang) by building a CRUD (Create, Read, Update, Delete) website using a MySQL database and Bootstrap 5. The project demonstrates the use of Golang's standard library without any external frameworks.

## Objectives

- Learn the fundamentals of web development using native Go.
- Establish a connection between Go and a MySQL database.
- Implement CRUD operations with SQL queries.
- Manage routing and HTTP requests using `http.HandlerFunc`.
- Integrate Bootstrap 5 for responsive front-end development.

## Project Overview

The project aims to create a simple website with basic CRUD functionalities for managing products and categories. It covers:

- Setting up a Go project.
- Connecting to a MySQL database.
- Implementing CRUD operations for products and categories.
- Handling web requests and routing using Go's standard library.
- Utilizing Bootstrap 5 for a modern and responsive UI.

## Prerequisites

Ensure you have the following installed before starting:

- [Go](https://golang.org/doc/install) (version 1.21.6 or higher recommended)
- [MySQL](https://dev.mysql.com/downloads/installer/)
- Basic knowledge of SQL, HTML, CSS, and JavaScript

## Getting Started

1. **Clone the Repository**

   Use the following command to clone the repository:

   ```bash
   git clone https://github.com/abdullahalhwyji/crud_web_golang.git
   cd crud_web_golang
   ```

2. **Install Go Dependencies**

   Run the following command to install the required Go dependencies:

   ```bash
   go mod tidy
   ```

3. **Set Up MySQL Database**

   - Create a MySQL database and run the following SQL scripts to create the required tables:

   ```sql
   CREATE TABLE `categories` (
     `id` int NOT NULL AUTO_INCREMENT,
     `name` varchar(100) NOT NULL,
     `created_at` timestamp NOT NULL,
     `updated_at` timestamp NOT NULL,
     PRIMARY KEY (`id`)
   ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

   CREATE TABLE `products` (
     `id` int NOT NULL AUTO_INCREMENT,
     `name` varchar(100) NOT NULL,
     `category_id` int NOT NULL,
     `stock` int NOT NULL,
     `description` text,
     `created_at` timestamp NOT NULL,
     `updated_at` timestamp NOT NULL,
     PRIMARY KEY (`id`)
   ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
   ```

   - Update the `config/database.go` file with your database credentials (username, password, database name).

4. **Run the Application**

   To start the Go server, execute:

   ```bash
   go run main.go
   ```

   The application will be available at `http://localhost:8000`.

5. **Open the Website**

   Visit `http://localhost:8000` in your web browser to access the website.

## Folder Structure

```
/crud_web_golang
│
├── /config/                # Configuration files
│   └── database.go         # Database connection setup
│
├── /controllers/           # Controllers for handling web requests
│   ├── categoriescontroller/
│   │   └── categoriescontroller.go
│   ├── homecontroller/
│   │   └── homecontroller.go
│   └── productcontroller/
│       └── productcontroller.go
│
├── /entities/              # Entity definitions
│   ├── category.go         # Category entity
│   └── product.go          # Product entity
│
├── /models/                # Data models for interacting with the database
│   ├── categoryModel/
│   │   └── categoryModel.go
│   └── productmodel/
│       └── productmodel.go
│
├── /views/                 # HTML templates for rendering web pages
│   ├── category/
│   │   ├── create.html
│   │   ├── edit.html
│   │   └── index.html
│   ├── home/
│   │   └── index.html
│   └── product/
│       ├── create.html
│       ├── detail.html
│       ├── edit.html
│       └── index.html
│
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies file
├── main.go                 # Main application entry point
└── README.md               # Project documentation
```

## Key Features

- **Database Connection:** Configured in `config/database.go` for connecting Go with MySQL.
- **Controllers:** Organized in the `controllers/` directory to handle different routes and functionalities.
- **Models:** Define the database interactions in `models/` directory, following a clean and maintainable structure.
- **Entities:** Represent the data structures used in the application.
- **Views:** HTML templates stored in `views/` directory for dynamic content rendering.

## Dependencies

This project uses the following Go modules:

- **[github.com/go-sql-driver/mysql v1.8.1](https://github.com/go-sql-driver/mysql)**: A MySQL driver for Go's `database/sql` package.
- **[filippo.io/edwards25519 v1.1.0](https://pkg.go.dev/filippo.io/edwards25519)**: A package providing an implementation of the Ed25519 digital signature algorithm. (Indirectly used by some other dependency)

### Go Module Setup

The `go.mod` file is already set up for you. If you clone the repository, make sure to install all dependencies by running:

```bash
go mod tidy
```

### go.mod File

Here’s what the `go.mod` file looks like:

```go
module github.com/abdullahalhwyji/crud_web_golang

go 1.21.6

require (
    filippo.io/edwards25519 v1.1.0 // indirect
    github.com/go-sql-driver/mysql v1.8.1 // indirect
)
```
