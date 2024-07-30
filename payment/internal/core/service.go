package core

// primary port

type PaymentService interface {
	CreatePayment(payment *Payment) error
	GetPaymentStatusByID(paymentID string) (*Payment, error)
}

type paymentServiceImpl struct {
	repo PaymentRepository
}

func NewPaymentService(repo PaymentRepository) PaymentService {
	return &paymentServiceImpl{repo: repo}
}

func (s *paymentServiceImpl) CreatePayment(payment *Payment) error {
	if err := s.repo.SavePayment(payment); err != nil {
		return err
	}
	return nil
}

func (s *paymentServiceImpl) GetPaymentStatusByID(paymentID string) (*Payment, error) {
	payment, err := s.repo.FindPaymentStatusByID(paymentID)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
