package transport

type RequestBook struct {
	Title   string `json:"title"`
	Year    int    `json:"year"`
	Author  string `json:"author"`
	Summary string `json:"summary"`
}

type RequestBookId struct {
	BookId string `json:"book_id"`
}