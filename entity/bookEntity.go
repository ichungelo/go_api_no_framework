package entity

type BookEntity struct {
	BookId  string `json:"book_id"` 
	Title   string `json:"title"`
	Year    int    `json:"year"`
	Author  string `json:"author"`
	Summary string `json:"summary" `
}