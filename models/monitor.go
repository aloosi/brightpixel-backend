package models

type Monitor struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageURL    string  `json:"imageUrl" db:"image_url"`
	Category    string  `json:"category" db:"category"`
	Brand       string  `json:"brand" db:"brand"`
	SizeGroup   string  `json:"sizeGroup" db:"size_group"`
}
