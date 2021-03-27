package notes

const (
	GetNotesListURL = "/api/notes/"
	PostNotesURL    = "/api/notes/"
	GetNotesURL     = "/api/notes/:id/"
	PutNotesURL     = "/api/notes/:id/"
	DeleteNoteURL   = "/api/notes/:id/"
)

// Note is a struct with a title and a body
type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Tags  string `json:"tags"`
}

type GetNotesListRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type PostNotesRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Tags  string `json:"tags"`
}

type PutNotesRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Tags  string `json:"tags"`
}

type DeleteNotesRequest struct{}

// NotesHypermedia holds the hypermedia format returned with GET, POST, PUT and DELETE calls
type NotesHypermedia struct {
	First string `json:"first"`
	Self  string `json:"self"`
	Last  string `json:"last"`
}

// GetNoteListResponse is returned when one or more notes match the given query parameters
type GetNoteListResponse struct {
	Links NotesHypermedia `json:"_links"`
	Notes []Note          `json:"notes"`
}

type GetNotesResponse struct {
	Note Note `json:"note"`
}

type PostNotesResponse struct {
	ID int `json:"id"`
}

type PutNotesResponse struct {
	ID int `json:"id"`
}

type DeleteNotesResponse struct {
	ID int `json:"id"`
}

// ValidationError is returned when the required input fails validation
type ErrorResponse struct {
	Error string `json:"error"`
}
