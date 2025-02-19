package core

//secondary port

type PaymentRepository interface {
	SavePayment(payment *Payment) error
	FindPaymentStatusByID(paymentID string) (*Payment, error)
	UpdatePaymentStatusByID(paymentID string, status string) error
}
