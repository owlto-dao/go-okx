package main

import (
	"log"

	"github.com/owlto-dao/go-okx/examples"
	"github.com/owlto-dao/go-okx/ws"
	"github.com/owlto-dao/go-okx/ws/private"
)

func main() {
	args := &ws.Args{
		InstType: "SPOT",
	}
	handler := func(c private.EventOrders) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeOrders(args, examples.Auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
