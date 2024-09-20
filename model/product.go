package model

type Product struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Price      float32 `json:"price" db:"price"`
	StockCount int     `json:"stock_count" db:"stock_count"`
}
