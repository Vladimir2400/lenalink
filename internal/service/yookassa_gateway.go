package service

import (
	"context"
	"fmt"

	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	yoorefund "github.com/rvinnie/yookassa-sdk-go/yookassa/refund"

	"github.com/lenalink/backend/internal/domain"
)

// YooKassaGateway implements PaymentGateway for YooKassa
type YooKassaGateway struct {
	paymentHandler *yookassa.PaymentHandler
	refundHandler  *yookassa.RefundHandler
	returnURL      string
}

// NewYooKassaGateway creates a new YooKassa payment gateway
func NewYooKassaGateway(shopID, secretKey, returnURL string) *YooKassaGateway {
	client := yookassa.NewClient(shopID, secretKey)
	return &YooKassaGateway{
		paymentHandler: yookassa.NewPaymentHandler(client),
		refundHandler:  yookassa.NewRefundHandler(client),
		returnURL:      returnURL,
	}
}

// ProcessPayment creates a payment in YooKassa
func (g *YooKassaGateway) ProcessPayment(ctx context.Context, payment *domain.Payment) error {
	// Build payment request
	paymentRequest := &yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    fmt.Sprintf("%.2f", payment.Amount),
			Currency: payment.Currency,
		},
		Confirmation: &yoopayment.Redirect{
			Type:      yoopayment.TypeRedirect,
			ReturnURL: g.returnURL,
		},
		Description: fmt.Sprintf("LenaLink: Бронирование %s", payment.OrderID),
		Metadata: map[string]string{
			"order_id":   payment.OrderID,
			"payment_id": payment.ID,
		},
		Capture: true, // Auto-capture payment
	}

	// Create payment in YooKassa
	resp, err := g.paymentHandler.CreatePayment(paymentRequest)
	if err != nil {
		return fmt.Errorf("yookassa create payment failed: %w", err)
	}

	// Save YooKassa payment ID
	payment.ProviderPaymentID = resp.ID

	// Save confirmation URL for redirect
	if resp.Confirmation != nil {
		if redirect, ok := resp.Confirmation.(*yoopayment.Redirect); ok {
			payment.ConfirmationURL = redirect.ConfirmationURL
		}
	}

	// Payment is pending until user completes it
	payment.Status = domain.PaymentPending

	return nil
}

// RefundPayment creates a refund in YooKassa
func (g *YooKassaGateway) RefundPayment(ctx context.Context, paymentID string, amount float64) error {
	refundRequest := &yoorefund.Refund{
		PaymentId: paymentID,
		Amount: &yoocommon.Amount{
			Value:    fmt.Sprintf("%.2f", amount),
			Currency: "RUB",
		},
		Description: "Возврат средств за отмену бронирования",
	}

	_, err := g.refundHandler.CreateRefund(refundRequest)
	if err != nil {
		return fmt.Errorf("yookassa refund failed: %w", err)
	}

	return nil
}

// GetPaymentStatus retrieves payment status from YooKassa
func (g *YooKassaGateway) GetPaymentStatus(ctx context.Context, paymentID string) (domain.PaymentStatus, error) {
	resp, err := g.paymentHandler.FindPayment(paymentID)
	if err != nil {
		return "", fmt.Errorf("yookassa get payment failed: %w", err)
	}

	// Map YooKassa status to domain status
	switch resp.Status {
	case yoopayment.Pending:
		return domain.PaymentPending, nil
	case yoopayment.WaitingForCapture:
		return domain.PaymentPending, nil
	case yoopayment.Succeeded:
		return domain.PaymentCompleted, nil
	case yoopayment.Canceled:
		return domain.PaymentFailed, nil
	default:
		return domain.PaymentPending, nil
	}
}
