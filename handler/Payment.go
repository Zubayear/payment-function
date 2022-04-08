package handler

import (
	"Payment/messages"
	"encoding/json"
	"github.com/google/uuid"
	"go-micro.dev/v4/broker"
	"time"

	log "go-micro.dev/v4/logger"
)

var (
	topicOrderPayment       = "order.topic.orderPayment"
	topicOrderPaymentUpdate = "order.topic.orderPaymentUpdate"
)

type Payment struct{}

func ConsumeMessage() {
	_, err := broker.Subscribe(topicOrderPayment, func(event broker.Event) error {
		var msg messages.OrderPaymentRequestMessage
		err := json.Unmarshal(event.Message().Body, &msg)
		if err != nil {
			return err
		}
		log.Infof("consume message, order payment request: %v", msg)
		// Call External api
		// Put the message in topic
		msgId, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		orderPaymentUpdateMsg := messages.NewOrderPaymentUpdateMessage(msgId, true)
		msgBody, err := json.Marshal(orderPaymentUpdateMsg)
		if err != nil {
			return err
		}
		msgToSend := &broker.Message{
			Header: map[string]string{
				"id":   msgId.String(),
				"time": time.Now().String(),
			},
			Body: msgBody,
		}
		err = broker.Publish(topicOrderPaymentUpdate, msgToSend)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}
