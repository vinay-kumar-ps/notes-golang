package main

import "github.com/gin-gonic/gin"

type Note struct {
	ID    int    `json:"id"`
	Title string `json"title"`
	Body  string `json"body"`
}

var notes = make(map[int]Note)
var currentID = 1

func createNote(c *gin.Context) {
	var newNote Note

	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newNote.ID = currentID
	notes[currentID] = newNote
	currentID++

	c.JSON(201, newNote)
}

func getNotes(c gin.Context) {
	var allNotes []Note
	for _, note := range notes {
		allNotes = append(allNotes, note)
	}
	c.JSON(200, allNotes)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api up",
		})
	})

	r.POST("/notes", createNote)
	r.GET("/notes", getNotes)
	r.Run(":8080")
}
