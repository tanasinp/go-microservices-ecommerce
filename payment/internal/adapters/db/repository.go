package adapters

import (
	"github.com/tanasinp/go-microservices-ecommerce/payment/internal/core"
	"gorm.io/gorm"
)

type gormPaymentRepository struct {
	db *gorm.DB
}

func NewGormPaymentRepository(db *gorm.DB) core.PaymentRepository {
	return &gormPaymentRepository{db: db}
}

func (r *gormPaymentRepository) SavePayment(payment *core.Payment) error {
	return r.db.Create(payment).Error
}

func (r *gormPaymentRepository) FindPaymentStatusByID(paymentID string) (*core.Payment, error) {
	var payment core.Payment
	if err := r.db.Where("id = ?", paymentID).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *gormPaymentRepository) UpdatePaymentStatusByID(paymentID string, status string) error {
	return r.db.Model(&core.Payment{}).Where("id = ?", paymentID).Update("status", status).Error
}
