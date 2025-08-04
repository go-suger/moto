package main

import (
	"log"

	"github.com/go-suger/moto"
)

type OrderState string

const (
	CREATED    OrderState = "created"
	PAID       OrderState = "paid"
	DELIVERING OrderState = "delivering"
	RECEIVED   OrderState = "received"
	DONE       OrderState = "done"
	CANCELLING OrderState = "cancelling"
	RETURNING  OrderState = "returning"
	CLOSED     OrderState = "closed"
)

type OrderEvent string

const (
	PAY     OrderEvent = "pay"
	DELIVER OrderEvent = "deliver"
	RECEIVE OrderEvent = "receive"
	CONFIRM OrderEvent = "confirm"

	CANCEL OrderEvent = "cancel"
	RETURN OrderEvent = "return"
	CLOSE  OrderEvent = "close"
)

type Order struct {
}

func main() {
	builder := moto.New[OrderState, OrderEvent, Order]()
	//
	builder.ExternalTransition().
		Form(CREATED).To(PAID).On(PAY).
		WhenFunc(func(order Order) bool {

			return false
		}).
		PerformFunc(func(from, to OrderState, event OrderEvent, order *Order) error {

			return nil
		})

	fsm, err := builder.Build()
	if err != nil {
		log.Println(err)
	}

	state, err := fsm.FireEvent(CREATED, PAY, &Order{})
	if err != nil {
		log.Println(err)
	}

	_ = state
}
