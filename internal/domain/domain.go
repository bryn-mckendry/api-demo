package domain

type Animal struct {
	Id      int
	Name    string `json:"name"`
	Species string `json:"species"`
}
