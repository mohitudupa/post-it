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

// GetNoteListHandler : GET method for notes API
func GetNoteListHandler(c *gin.Context) {
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

// PostNotesHandler : POST method for notes API
func PostNotesHandler(c *gin.Context) {
	req := PostNotesRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
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

	id, err := DB.Add(n)
	if err != nil {
		// Return a 500 internal server error
		e := ErrorResponse{Error: err.Error()}
		c.JSON(500, e)
		return
	}
	resp := PostNotesResponse{ID: id}
	c.JSON(201, resp)
}

// GetNotesHandler : Get method for notes API
func GetNotesHandler(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: "id must be an integer"}
		c.JSON(400, e)
		return
	}

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
	resp := GetNotesResponse{Note: n}
	c.JSON(200, resp)
}

// PutNotesHandler : PUT method for notes API
func PutNotesHandler(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		// Return a 400 bad request
		e := ErrorResponse{Error: "id must be an integer"}
		c.JSON(400, e)
		return
	}

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

	resp := PutNotesResponse{ID: id}
	c.JSON(201, resp)
}

// DeleteNoteHandler : DELETE method for notes API
func DeleteNoteHandler(c *gin.Context) {
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
