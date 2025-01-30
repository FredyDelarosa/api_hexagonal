package entities

type IDE struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Developer   string `json:"developer"`
	ReleaseYear int    `json:"release_year"`
}
