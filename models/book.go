package models

var Books = []Book{
	{ID: "1", Name: "Atomic Habits", Author: "James Clear"},
}

// Book Type
type Book struct {
	ID     string `json :"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
