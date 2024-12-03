package main

import (
	"log"

	"github.com/owlto-dao/go-okx/examples/rest"
	"github.com/owlto-dao/go-okx/rest/api"
	"github.com/owlto-dao/go-okx/rest/api/account"
)

func main() {
	param := &account.PostPositionMarginBalanceParam{
		InstId:  "BTC-USDT-SWAP",
		PosSide: api.PosSideNet,
		Type:    api.TypeReduce,
		Amt:     "1",
	}
	req, resp := account.NewPostPositionMarginBalance(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.PostPositionMarginBalanceResponse))
}
