<h1 align="center">POS APP RESTfull API Golang and PostgreSQL</h1>
<p align="center">
  <a href="https://www.fazztrack.com/">
    <img src="https://www.fazztrack.com/_nuxt/img/fazztrack-logo-color.db4c9cc.svg" width="200px" alt="" />
  </a>
</p>
<hr/>


## Built With

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

2. Installing Package

   ```bash
   go get -u github.com/<package_name>
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




