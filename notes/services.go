package notes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	// NotesURL is the url for notes API
	NotesURL      = "/api/notes"
	defaultOffset = 0
	defaultLimit  = 100
)

// GetNoteListHandler Gin HTTP handler for getting a list of notes
func GetNoteListHandler(c *gin.Context) {
	// swagger:route GET /notes notes listNotes
	//
	//
	// This will return notes based on the offset and limit values
	// Limit ranges from 1 - 500
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: GetNoteListResponse
	//       400: ErrorResponse
	//       500: ErrorResponse

	// Extract query parameters
	qry := GetNotesListRequest{defaultOffset, defaultLimit}
	if err := c.ShouldBindQuery(&qry); err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: err.Error()}
		c.JSON(400, e)
		return
	}

	offset := int(qry.Offset)
	limit := int(qry.Limit)

	// Validate limit constraints
	if !(1 <= limit && limit <= 500) {
		// Return a 400 bad request
		e := ErrorResponse{Error: "limit must be between 1 and 500"}
		c.JSON(400, e)
		return
	}

	// Get notes
	ns, err := DB.List(offset, limit)
	if err != nil {
		// Return a 500 internal server error
		e := ErrorResponse{Error: err.Error()}
		c.JSON(500, e)
		return
	}

	// Get count of notes
	count, err := DB.Count()
	if err != nil {
		// Return a 500 internal server error
		e := ErrorResponse{Error: err.Error()}
		c.JSON(500, e)
		return
	}

	// Return response
	resp := GetNoteListResponse{
		Links: GetLinks(offset, limit, count),
		Notes: ns,
	}
	c.JSON(200, resp)
}

// PostNotesHandler Gin HTTP handler for creating notes
func PostNotesHandler(c *gin.Context) {
	// swagger:route POST /notes notes createNotes
	//
	//
	// This will create a note with a title, body and a list of tags
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       201: PostNotesResponse
	//       400: ErrorResponse
	//       500: ErrorResponse

	// Read POST request body
	req := PostNotesRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: err.Error()}
		c.JSON(400, e)
		return
	}

	// Add note to DB
	n := Note{
		Title: req.Title,
		Body:  req.Body,
		Tags:  req.Tags,
	}
	id, err := DB.Add(n)
	if err != nil {
		// Return a 500 internal server error
		e := ErrorResponse{Error: err.Error()}
		c.JSON(500, e)
		return
	}

	// Return new note ID
	resp := PostNotesResponse{ID: id}
	c.JSON(201, resp)
}

// GetNotesHandler Gin HTTP handler for getting a note by it's ID
func GetNotesHandler(c *gin.Context) {
	// swagger:route GET /notes/{noteId}/ notes getNote
	//
	//
	// This will return a note given it's note ID
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: GetNotesResponse
	//       400: ErrorResponse
	//       404: ErrorResponse
	//       500: ErrorResponse

	// Get Note ID from URL
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: "id must be an integer"}
		c.JSON(400, e)
		return
	}

	// Look for note in the DB
	n, err := DB.Get(id)
	if err != nil && err.Error() == "404 Not Found" {
		// Return a 404 not found
		e := ErrorResponse{Error: err.Error()}
		c.JSON(404, e)
		return
	} else if err != nil {
		// Return a 500 internal server error
		e := ErrorResponse{Error: err.Error()}
		c.JSON(500, e)
		return
	}

	// Return Note
	resp := GetNotesResponse{Note: n}
	c.JSON(200, resp)
}

// PutNotesHandler Gin HTTP handler for updating notes
func PutNotesHandler(c *gin.Context) {
	// swagger:route PUT /notes/{noteId}/ notes updateNote
	//
	//
	// This will update(replace) a note given it's note ID
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       201: PutNotesResponse
	//       400: ErrorResponse
	//       404: ErrorResponse
	//       500: ErrorResponse

	// Get Note ID from URL
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: "id must be an integer"}
		c.JSON(400, e)
		return
	}

	// Read PUT request body
	req := PutNotesRequest{}
	if err = c.ShouldBindJSON(&req); err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: err.Error()}
		c.JSON(400, e)
		return
	}

	n := Note{
		Title: req.Title,
		Body:  req.Body,
		Tags:  req.Tags,
	}

	// Update note on DB
	err = DB.Update(id, n)
	if err != nil && err.Error() == "404 Not Found" {
		// Return a 404 not found
		e := ErrorResponse{Error: err.Error()}
		c.JSON(404, e)
		return
	} else if err != nil {
		// Return a 500 internal server error
		e := ErrorResponse{Error: err.Error()}
		c.JSON(500, e)
		return
	}

	// Return ID of updated note
	resp := PutNotesResponse{ID: id}
	c.JSON(201, resp)
}

// DeleteNoteHandler Gin HTTP handler for deleting notes
func DeleteNoteHandler(c *gin.Context) {
	// swagger:route DELETE /notes/{noteId}/ notes deleteNote
	//
	//
	// This will delete a note given it's note ID
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: DeleteNotesResponse
	//       400: ErrorResponse
	//       404: ErrorResponse
	//       500: ErrorResponse

	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: "id must be an integer"}
		c.JSON(400, e)
		return
	}

	err = DB.Delete(id)
	if err != nil {
		// Return a 404 not found
		e := ErrorResponse{Error: err.Error()}
		c.JSON(404, e)
	}
	resp := DeleteNotesResponse{ID: id}
	c.JSON(200, resp)
}

// GetLinks fetches next, self and last links for a resource
func GetLinks(o int, l int, c int) NotesHypermedia {
	lp := int(c / l)
	ln := NotesHypermedia{
		First: fmt.Sprintf("%s?offset=%d&limit=%d", NotesURL, 0, l),
		Self:  fmt.Sprintf("%s?offset=%d&limit=%d", NotesURL, o, l),
		Last:  fmt.Sprintf("%s?offset=%d&limit=%d", NotesURL, lp, l),
	}

	return ln
}
