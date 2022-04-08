package messages

import "github.com/google/uuid"

type OrderPaymentRequestMessage struct {
	OrderId                              uuid.UUID
	Total                                float64
	CardNumber, CardName, CardExpiration string
}

func NewOrderPaymentRequestMessage(orderId uuid.UUID, total float64, cardNumber string, cardName string, cardExpiration string) *OrderPaymentRequestMessage {
	return &OrderPaymentRequestMessage{OrderId: orderId, Total: total, CardNumber: cardNumber, CardName: cardName, CardExpiration: cardExpiration}
}
