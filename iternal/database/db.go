package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// AddUrlDB inserts a new URL into the database.
//
// Parameters:
// - hash: a string representing the hash of the URL.
// - longUrl: a string representing the long URL.
//
// Return type: None.
func AddUrlDB(hash string, longUrl string) {
	connStr := "user=postgres password=1 dbname=test3 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO urls (hash, longurl) VALUES ($1, $2)", hash, longUrl)
	if err != nil {
		panic(err)
	}
}

// CheckUrlDB checks if a given URL exists in the database.
//
// Parameters:
// - rawUrl: a string representing the URL to be checked.
//
// Returns:
// - a string representing the shortened URL if the URL exists in the database,
//   or an empty string if the URL is not found.
func CheckUrlDB(rawUrl string) string {
	var url string
	connStr := "user=postgres password=1 dbname=test3 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.QueryRow("SELECT hash FROM urls WHERE longurl = $1", rawUrl).Scan(&url)
	if err == sql.ErrNoRows {
		return ""
	}
	return "http://localhost:8080/" + url
}
