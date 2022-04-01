package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student represents data about a record Student.
type album struct {
	StudentID string  `json:"StudentID"`
	Name      string  `json:"Name"`
	Address   string  `json:"Address"`
	Marks     float64 `json:"marks "`
}

// student slice to seed record album data.
var albums = []album{
	{StudentID: "1", Name: "Pooja", Address: "Pune", Marks: 70.99},
	{StudentID: "2", Name: "Khayti", Address: "Mumbai", Marks: 85.99},
	{StudentID: "3", Name: "Arth", Address: "Rajasthan", Marks: 75.99},
}

func main() {
	router := gin.Default()
	router.GET("/Student", getStudent)
	router.GET("/Student/:id", getAlbumByID)
	router.POST("/Student", postStudent)

	router.Run("localhost:8080")
}

// getstudent responds with the list of all student as JSON.
func getStudent(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// poststudent adds an student from JSON received in the request body.
func postStudent(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.StudentID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
