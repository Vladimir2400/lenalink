package service

import (
	"context"
	"fmt"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/pkg/utils"
)

// PaymentGateway defines interface for payment processing
type PaymentGateway interface {
	ProcessPayment(ctx context.Context, payment *domain.Payment) error
	RefundPayment(ctx context.Context, paymentID string, amount float64) error
	GetPaymentStatus(ctx context.Context, paymentID string) (domain.PaymentStatus, error)
}

// PaymentService handles payment processing
type PaymentService struct {
	gateway PaymentGateway
}

// NewPaymentService creates a new payment service
func NewPaymentService(gateway PaymentGateway) *PaymentService {
	return &PaymentService{gateway: gateway}
}

// CreatePayment creates a new payment for booking
func (ps *PaymentService) CreatePayment(orderID string, amount float64, method domain.PaymentMethod) *domain.Payment {
	return &domain.Payment{
		ID:        utils.GenerateID(),
		OrderID:   orderID,
		Amount:    amount,
		Currency:  "RUB",
		Method:    method,
		Status:    domain.PaymentPending,
		CreatedAt: time.Now(),
	}
}

// ProcessPayment processes a payment
func (ps *PaymentService) ProcessPayment(ctx context.Context, payment *domain.Payment) error {
	// Process payment through gateway
	if err := ps.gateway.ProcessPayment(ctx, payment); err != nil {
		payment.Status = domain.PaymentFailed
		payment.FailureReason = err.Error()
		return fmt.Errorf("payment processing failed: %w", err)
	}

	payment.Status = domain.PaymentCompleted
	now := time.Now()
	payment.CompletedAt = &now
	return nil
}

// RefundPayment processes a refund
func (ps *PaymentService) RefundPayment(ctx context.Context, payment *domain.Payment) error {
	if payment.Status != domain.PaymentCompleted {
		return fmt.Errorf("cannot refund payment in status: %s", payment.Status)
	}

	if err := ps.gateway.RefundPayment(ctx, payment.ID, payment.Amount); err != nil {
		return fmt.Errorf("refund failed: %w", err)
	}

	payment.Status = domain.PaymentRefunded
	return nil
}

// CheckPaymentStatus checks payment status from gateway
func (ps *PaymentService) CheckPaymentStatus(ctx context.Context, payment *domain.Payment) (domain.PaymentStatus, error) {
	return ps.gateway.GetPaymentStatus(ctx, payment.ID)
}

// --- Mock Payment Gateway for MVP/Hackathon ---

// MockPaymentGateway simulates payment processing for testing
type MockPaymentGateway struct {
	failureRate float64 // Probability of payment failure (0.0 - 1.0)
}

// NewMockPaymentGateway creates a mock payment gateway
func NewMockPaymentGateway(failureRate float64) *MockPaymentGateway {
	return &MockPaymentGateway{
		failureRate: failureRate,
	}
}

// ProcessPayment simulates payment processing
func (mpg *MockPaymentGateway) ProcessPayment(ctx context.Context, payment *domain.Payment) error {
	// Simulate processing delay
	time.Sleep(100 * time.Millisecond)

	// Generate mock provider payment ID
	payment.ProviderPaymentID = fmt.Sprintf("MOCK-PAY-%s", utils.GenerateID()[:8])

	// Simulate random failures for testing
	// For production, this would be replaced with real payment gateway integration
	// (uncomment for testing failure scenarios)
	// if rand.Float64() < mpg.failureRate {
	// 	return fmt.Errorf("mock payment failed: insufficient funds")
	// }

	// Always succeed in mock mode for hackathon demo
	return nil
}

// RefundPayment simulates refund processing
func (mpg *MockPaymentGateway) RefundPayment(ctx context.Context, paymentID string, amount float64) error {
	// Simulate processing delay
	time.Sleep(100 * time.Millisecond)

	// Always succeed in mock mode
	return nil
}

// GetPaymentStatus returns mock payment status
func (mpg *MockPaymentGateway) GetPaymentStatus(ctx context.Context, paymentID string) (domain.PaymentStatus, error) {
	// In mock mode, assume all payments are completed
	return domain.PaymentCompleted, nil
}

// --- Future: Real Payment Gateway Implementations ---

// YooKassaGateway would implement real YooKassa integration
// type YooKassaGateway struct {
// 	apiKey    string
// 	shopID    string
// 	apiClient *http.Client
// }

// CloudPaymentsGateway would implement real CloudPayments integration
// type CloudPaymentsGateway struct {
// 	publicID  string
// 	apiSecret string
// 	apiClient *http.Client
// }
