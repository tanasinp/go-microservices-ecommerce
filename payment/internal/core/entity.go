package core

type Payment struct {
	ID      string `gorm:"primaryKey"`
	OrderID string
	UserID  string
	Status  string
	Total   float64
}
