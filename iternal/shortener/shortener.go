package shortener

import (
	"flag"
	"net/http"
	"time"

	"shortener-api/iternal/database"

	"github.com/gin-gonic/gin"

	"github.com/sqids/sqids-go"
)

type UrlsStorage struct {
	urls map[string]Url
}

type Url struct {
	Hash    string
	LongUrl string
}

// addUrl adds a URL to the storage.
//
// The URL is added with its Hash as the key.
// No return value.
func (s UrlsStorage) addUrl(url Url) {
	s.urls[url.Hash] = url
}

// getUrl retrieves the original URL from the storage by its hash.
//
// hash - the shortened URL hash.
// string - the original URL.
func (s UrlsStorage) getUrl(hash string) string {
	url := s.urls[hash]
	return url.LongUrl
}

// Storage returns a new instance of UrlsStorage.
func Storage() *UrlsStorage {
	var s UrlsStorage
	s.urls = make(map[string]Url)
	return &s
}

var storage = Storage()
var useDB = flag.Bool("d", false, "save url into database")

// CreateShortUrl is the handler for the POST / route.
func CreateShortUrl(c *gin.Context) {
	type RequestBody struct {
		Body string // The URL to be shortened
	}
	var rawUrl RequestBody
	if err := c.BindJSON(&rawUrl); err != nil {
		// Return a 400 Bad Request if the request body is invalid
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	checkedUrl := checkUrl(rawUrl.Body, *storage)
	if checkedUrl != "" {
		// Return a 200 OK if the URL is already stored in memory
		c.IndentedJSON(http.StatusOK, checkedUrl)
	} else {
		hash := createEndpoint()
		url := Url{Hash: hash, LongUrl: rawUrl.Body}
		flag.Parse()
		if *useDB {
			// Store the URL in the database if the flag is set
			database.AddUrlDB(url.Hash, url.LongUrl)
		}
		UrlsStorage.addUrl(*storage, url)
		// Return a 201 Created with the shortened URL
		c.IndentedJSON(http.StatusCreated, "http://localhost:8080/"+hash)
	}
}

// GetRawUrl retrieves the original URL from the storage by its hash.
//
// c - the gin context.
// Returns the original URL as a JSON response.
func GetRawUrl(c *gin.Context) {
	hash := c.Param("url")
	rawUrl := UrlsStorage.getUrl(*storage, hash)
	if rawUrl == "" {
		c.IndentedJSON(http.StatusNotFound, "Not found")
	} else {
		c.IndentedJSON(http.StatusFound, rawUrl)
	}

}

// checkUrl checks if a given URL is already stored in memory or in the database.
//
// rawUrl - the URL to be checked.
// s - the storage containing the URLs.
// Returns the shortened URL if found, otherwise an empty string.
func checkUrl(rawUrl string, s UrlsStorage) string {
	flag.Parse()
	if *useDB {
		return database.CheckUrlDB(rawUrl)
	}
	for hash, url := range s.urls {
		if rawUrl == url.LongUrl {
			return "http://localhost:8080/" + hash
		}
	}
	return ""
}

// createEndpoint generates a unique endpoint identifier.
//
// No parameters.
// Returns a string representing the endpoint identifier.
func createEndpoint() string {
	s, _ := sqids.New()
	now := time.Now()
	id, _ := s.Encode([]uint64{uint64(now.Unix())})
	return id
}
