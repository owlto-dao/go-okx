package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := []*trade.PostCancelOrderParam{
		&trade.PostCancelOrderParam{
			InstId: "OKB-USDT",
			OrdId:  "123456789",
		},
		&trade.PostCancelOrderParam{
			InstId: "XRP-USDT",
			OrdId:  "123456789",
		},
	}
	req, resp := trade.NewPostCancelBatchOrders(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostCancelOrderResponse))
}