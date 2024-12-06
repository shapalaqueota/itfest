package models

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	ImageURL    string `json:"image"`
}
