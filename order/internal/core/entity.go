package core

type Order struct {
	ID string `gorm:"primaryKey"`
	// gorm.Model
	UserID  string
	Items   []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	Address string
	Total   float64
	Status  string
}

type OrderItem struct {
	ID string `gorm:"primaryKey"`
	// gorm.Model
	ProductID string
	Quantity  int
	Price     float64
	OrderID   string
}
