package domain

type Product struct {
	ID            uint    `json:"id"`
	ProductName   string  `json:"product_name"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	CreatedAt     int64   `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt     int64   `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"updated_at"`
}
