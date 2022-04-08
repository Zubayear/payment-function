package messages

import "github.com/google/uuid"

type OrderPaymentUpdateMessage struct {
	OrderId       uuid.UUID
	PaymentStatus bool
}

func NewOrderPaymentUpdateMessage(orderId uuid.UUID, paymentStatus bool) *OrderPaymentUpdateMessage {
	return &OrderPaymentUpdateMessage{OrderId: orderId, PaymentStatus: paymentStatus}
}
