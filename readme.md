# URL Shortener Project
==========================
Проект выполнен в рамках тестового задания на стажировку.

### Overview
***
This is a simple URL shortener project written in Go. It allows users to shorten long URLs and store them in a database.

### Features
***
- Shorten long URLs to a shorter format
- Store shortened URLs in a PostgreSQL database
- Retrieve original URL from shortened URL
### Dependencies
***
- Go 1.14 or higher
- PostgreSQL 12 or higher
- Gin framework for routing
- sql package for database interactions
- squid  package generate unique IDs from numbers
### Installation
***
#### Prerequisites
- Install Go 1.14 or higher from the official Go website:
  https://golang.org/doc/install
- Install PostgreSQL 12 or higher from the official PostgreSQL website:
https://www.postgresql.org/download/
- Install the required Go packages using the following command:
```bash
go get -u github.com/gin-gonic/gin 
go get -u github.com/lib/pq
go get github.com/sqids/sqids-go
```

#### Project Setup

*   Clone the project repository using the following command:

```bash
git clone https://github.com/DenisLukianov21/swoyo-testovoe.git
```
* Navigate to the project directory:
```bash
cd swoyo_testovoe
```

*   Build the project using the following command:

```bash
go build main.go
```
### Configuration
***
##### Database Configuration
Update the connStr variable in the CheckUrlDB function to match your PostgreSQL database connection string.
##### Server Configuration
Update the router.Run function to change the server port or address.
### Usage
***
##### Shortening URLs
* Send a POST request to the / endpoint with the long URL as the request body.
* The shortened URL will be returned in the response.
##### Retrieving Original URLs
* Send a GET request to the /:url endpoint with the shortened URL as the path parameter.
* The original URL will be returned in the response.
### API Documentation
***
###### POST /
- Shorten a long URL.
- Request Body: longUrl (string)
- Response: shortUrl (string)
###### GET /:url
- Retrieve the original URL from a shortened URL.
- Path Parameter: url (string)
- Response: originalUrl (string)
