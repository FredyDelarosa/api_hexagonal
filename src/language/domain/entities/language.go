package entities

type Language struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Paradigm    string `json:"paradigm"`
	CreatedYear int    `json:"created_year"`
}
