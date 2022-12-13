package response

import "time"

type ProductCreateResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_Id"`
	CreatedAt  time.Time `json:"created_at"`
}
