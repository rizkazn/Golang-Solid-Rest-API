<h1 align="center">POS APP RESTfull API Golang and PostgreSQL</h1>
<p align="center">
  <a href="https://www.fazztrack.com/">
    <img src="https://www.fazztrack.com/_nuxt/img/fazztrack-logo-color.cba88b7.svg" width="200px" alt="Fazztracklogo.svg" />
  </a>
</p>
<hr/>

## Architecture 
<p align="center">
  <img src="https://res.cloudinary.com/rizkazn/image/upload/v1639242707/Product/Go_Arch_u4lanx.png" widht="700px" alt="" /> 
</p>

<!-- ABOUT THE PROJECT -->
## About The Project

Point Of Sale (POS) APP is designed to help small businesses with keeping track of customers, items and inventory, and generate reports based on sales. This website uses Golang and PostgreSQL as the data backend.	

## Built With

[![Golang](https://img.shields.io/badge/Golang-4.x-blue.svg?style=rounded-square)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v.13.3-blue.svg?style=rounded-square)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-v.6.2-red.svg?style=rounded-square)](https://redis.io/)

## Feature
- CRUD
- Error Handling
- Authentication and Authorization
- Search and Filter
- Middleware
- Payload Validation
- Hash Password & JWT
- DB Migrate
- Integration Testing & Assertion
- Mocking
- Solid Principle

## Installation Steps

1. Clone the repository

   ```bash
    https://github.com/rizkazn/Golang-Solid-Rest-API.git
    ```

2. Install Dependencies

   ```bash
   go build
   ```

3. Add .env file at root folder project

   ```sh
    host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
   ```

4. Test the app

   ```bash
   go test
   ```


5. Run the app

   ```bash
   go run main.go
   ```

6. You are all set!

   ```bash
   Backend running at: http://localhost:8080/api/v1
   ```





