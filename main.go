package main

import (
	"postit/notes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET(notes.GetNotesListURL, notes.GetNoteListHandler)
	r.POST(notes.PostNotesURL, notes.PostNotesHandler)
	r.GET(notes.GetNotesURL, notes.GetNotesHandler)
	r.PUT(notes.PutNotesURL, notes.PutNotesHandler)
	r.DELETE(notes.DeleteNoteURL, notes.DeleteNoteHandler)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}
