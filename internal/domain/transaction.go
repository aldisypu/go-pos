package domain

type Transaction struct {
	ID          uint  `json:"id"`
	TotalAmount int   `json:"total_amount"`
	CreatedAt   int64 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   int64 `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"updated_at"`
}
