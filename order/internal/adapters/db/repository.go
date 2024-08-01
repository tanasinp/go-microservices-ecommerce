package adapters

//database access
import (
	"github.com/tanasinp/go-microservices-ecommerce/order/internal/core"
	"gorm.io/gorm"
)

type gormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &gormOrderRepository{db: db}
}

func (r *gormOrderRepository) Save(order *core.Order) error {
	return r.db.Create(order).Error
}

func (r *gormOrderRepository) FindByID(id string) (*core.Order, error) {
	var order core.Order
	if err := r.db.Preload("Items").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *gormOrderRepository) UpdateOrderStatusByID(orderID string, status string) error {
	return r.db.Model(&core.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
