package main

import (
	"context"
	"fmt"
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
	Id string
}

func main() {
	builder := moto.New[OrderState, OrderEvent, Order]()
	//
	builder.ExternalTransition().
		Form(CREATED, DONE).To(PAID).On(PAY).
		WhenFunc(func(ctx context.Context, order Order) bool {
			return true
		}).
		PerformFunc(func(ctx context.Context, from, to OrderState, event OrderEvent, order *Order) error {
			log.Println("dad")
			order.Id = "3dadd"
			return nil
		})

	builder.ExternalTransition().
		Form(CREATED, DONE).To(PAID).On(DELIVER).
		WhenFunc(func(ctx context.Context, context Order) bool {
			return true
		}).
		PerformFunc(func(ctx context.Context, from, to OrderState, event OrderEvent, order *Order) error {
			order.Id = "3dadd"
			return nil
		})
	//
	fsm, err := builder.Build()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(fsm.GenerateMermaidGraph())

	order := &Order{
		Id: "1",
	}
	fmt.Println(order.Id)

	state, err := fsm.FireEvent(context.Background(), DONE, PAY, order)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(state)
	fmt.Println(order.Id)
}
