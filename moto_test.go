package moto

import (
	"context"
	"testing"
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

func TestMoto(t *testing.T) {
	builder := New[OrderState, OrderEvent, Order]()
	//
	builder.ExternalTransition().
		Form(CREATED).To(PAID).On(PAY).
		When(func(form, to OrderState, context Order) {

		}).
		Perform(func(form, to OrderState, context Order) {

		})

	fsm, err := builder.Build(context.Background())
	if err != nil {
		t.Error(err)
	}

	if err = fsm.FireEvent(CREATED, PAY, Order{}); err != nil {
		t.Error(err)
	}
}
