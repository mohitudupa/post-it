package notes

const (
	GetNotesListURL = "/api/notes/"
	PostNotesURL    = "/api/notes/"
	GetNotesURL     = "/api/notes/:id/"
	PutNotesURL     = "/api/notes/:id/"
	DeleteNoteURL   = "/api/notes/:id/"
)

// Note is the struct type for the note
// swagger:model Note
type Note struct {
	// ID of the note
	ID int `json:"id"`
	// Title of the note
	Title string `json:"title"`
	// Body of the note
	Body string `json:"body"`
	// Tags for the note. Usually space or comma separated list of things
	Tags string `json:"tags"`
}

// GetNotesListRequest is the struct type for GET query params of a listNotes request
// swagger:model GetNotesListRequest
type GetNotesListRequest struct {
	// Offset is the start index of the listNotes page. minValue: 0
	Offset int `form:"offset"`
	// Limit is the number of records to return per request. minValue 0. MaxValue 500
	Limit int `form:"limit"`
}

// PostNotesRequest is the struct type for POST body of a createNotes request
// swagger:model PostNotesRequest
type PostNotesRequest struct {
	// Title of the note
	Title string `json:"title"`
	// Body of the note
	Body string `json:"body"`
	// Tags for the note. Usually space or comma separated list of things
	Tags string `json:"tags"`
}

// PutNotesRequest is the struct type for POST body of a updateNotes request
// swagger:model PutNotesRequest
type PutNotesRequest struct {
	// Title of the note
	Title string `json:"title"`
	// Body of the note
	Body string `json:"body"`
	// Tags for the note. Usually space or comma separated list of things
	Tags string `json:"tags"`
}

// DeleteNotesRequest is the struct type for DELETE body of a deleteNotes request
// swagger:model DeleteNotesRequest
type DeleteNotesRequest struct{}

// NotesHypermedia is the struct type for links in return responses for listNotes
// swagger:model NotesHypermedia
type NotesHypermedia struct {
	// First is the link for the first set of notes
	First string `json:"first"`
	// Self is the link for the current set of notes
	Self string `json:"self"`
	// Last is the link for the last set of notes
	Last string `json:"last"`
}

// GetNoteListResponse is the struct type for return responses of listNotes
// swagger:response GetNoteListResponse
type GetNoteListResponse struct {
	// Links holds the links to the first, current and the last set of records
	Links NotesHypermedia `json:"_links"`
	// Notes is the list of notes
	Notes []Note `json:"notes"`
}

// GetNotesResponse is the struct type for return responses of getNotes
// swagger:response GetNotesResponse
type GetNotesResponse struct {
	// Note holds the ID, title, body and tags of the requested note
	Note Note `json:"note"`
}

// PostNotesResponse is the struct type for return responses of createNotes
// swagger:response PostNotesResponse
type PostNotesResponse struct {
	// ID of the newly created note
	ID int `json:"id"`
}

// PutNotesResponse is the struct type for return responses of updateNotes
// swagger:response PutNotesResponse
type PutNotesResponse struct {
	// ID of the updated note
	ID int `json:"id"`
}

// DeleteNotesResponse is the struct type for return responses of deleteNotes
// swagger:response DeleteNotesResponse
type DeleteNotesResponse struct {
	// ID of the deleted note
	ID int `json:"id"`
}

// ErrorResponse is the struct type for return responses of 4xx/5xx errors
// swagger:response ErrorResponse
type ErrorResponse struct {
	// Error message
	Error string `json:"error"`
}
