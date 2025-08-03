package main

import (
	"context"
	"log"
	"moto"
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
		When(func(form, to OrderState, context Order) {

		}).
		Perform(func(form, to OrderState, context Order) {

		})

	fsm, err := builder.Build(context.Background())
	if err != nil {
		log.Println(err)
	}

	if err = fsm.FireEvent(CREATED, PAY, Order{}); err != nil {
		log.Println(err)
	}
}
