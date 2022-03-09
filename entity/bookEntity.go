package entity

type BookEntity struct {
	BookId  string `json:"book_id" binding:"required"` 
	Title   string `json:"title" binding:"required"`
	Year    int    `json:"year" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Summary string `json:"summary" `
}