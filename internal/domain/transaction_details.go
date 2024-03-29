package domain

type TransactionDetail struct {
	ID            uint  `json:"id"`
	TransactionID uint  `json:"transaction_id"`
	ProductID     uint  `json:"product_id"`
	Quantity      int   `json:"quantity"`
	Subtotal      int   `json:"subtotal"`
	CreatedAt     int64 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt     int64 `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"updated_at"`

	Transaction Transaction `gorm:"foreignKey:TransactionID;references:id" json:"-"`
	Product     Product     `gorm:"foreignKey:ProductID;references:id" json:"-"`
}
