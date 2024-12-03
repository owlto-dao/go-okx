package main

import (
	"log"

	"github.com/owlto-dao/go-okx/ws"
	"github.com/owlto-dao/go-okx/ws/public"
)

func main() {
	args := &ws.Args{
		Channel: "mark-price-candle1m",
		InstId:  "BTC-USDT",
	}
	handler := func(c public.EventMarkPriceKline) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeMarkPriceKline(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
