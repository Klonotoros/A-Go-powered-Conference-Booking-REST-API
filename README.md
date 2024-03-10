# A-Go-powered-Conference-Booking-REST-API

## Overview

Welcome to the A-Go-powered-Conference-Booking-REST-API project! This API is built with Go (Golang) and utilizes the GIN framework, SQL for data storage, and JWT for token-based authentication. Below, you'll find a brief guide on the API, its endpoints, and technologies used.

## Installation

To get started, make sure you have Go installed on your machine. Clone the repository and run the following commands:

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/go-sql-driver/mysql
```

## Endpoints

- 📅 **GET /conferences**: Retrieve a list of available conferences.
- 📆 **GET /conferences/<id>**: Get details of a specific event.
- 📌 **POST /conferences**: Create a new bookable event (Authentication required).
- 🔄 **PUT /conferences/<id>**: Update an existing event (Authentication required/Only by creator).
- 🗑️ **DELETE /conferences/<id>**: Delete an event (Authentication required/Only by creator).
- 📌 **POST /signup**: Create a new user.
- 📌 **POST /login**: Authenticate user and receive JWT token.
- 📌 **POST /conferences/<id>/register**: Register user for a specific event (Authentication required).
- 🗑️ **DELETE /conferences/<id>/register**: Cancel registration for a specific event (Authentication required).

## Database schema

![image](https://github.com/Klonotoros/A-Go-powered-Conference-Booking-REST-API/assets/7630626/46f2cb62-b773-4971-82c7-4f64fca3ad58)

*"Don't panic when you see the 'registrations' table; it's just an example."*

## Technologies Used

- 🐹 **Go (Golang)**: Efficient and performant language for backend development.
- 🍸 **GIN Framework**: Lightweight HTTP framework for building APIs.
- 🗃️ **SQL (Database)**: Utilized for persistent data storage.
- 🔐 **JWT (JSON Web Tokens)**: Token-based authentication for secure communication.
- 📬 **Postman**: API development and testing tool.
- 💻 **IntelliJ GoLand**: Integrated development environment for Go.

## Quick Start

1. Clone the repository.
2. Install dependencies as mentioned in the Installation section.
3. Configure the database connection and JWT secret key in the appropriate configuration files.
4. Run the application using `go run main.go`.
5. Explore and interact with the API using the documented endpoints.

Feel free to reach out if you have any questions or if you'd like to discuss this project further. Happy coding!
