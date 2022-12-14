package entity

import "time"

type Transaction struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	ProductID  int       `json:"product_id"`
	Product    Product   `json:"-"`
	UserID     int       `json:"user_id"`
	User       User      `json:"-"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Transaction) TableName() string {
	return "transaction_history"
}
