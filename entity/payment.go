package entity

type Payment struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Amount    float64 `gorm:"not null"`
	Currency  string  `gorm:"size:3;not null"`
	Status    string  `gorm:"size:20;not null"`
	CreatedAt int64   `gorm:"autoCreateTime"`
	UpdatedAt int64   `gorm:"autoUpdateTime"`
}
