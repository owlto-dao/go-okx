package main

import (
	"log"

	"github.com/owlto-dao/go-okx/examples/rest"
	"github.com/owlto-dao/go-okx/rest/api"
	"github.com/owlto-dao/go-okx/rest/api/account"
)

func main() {
	param := &account.GetMaxAvailSizeParam{
		InstId: "BTC-USDT",
		TdMode: api.TdModeCross,
		Ccy:    "USDT",
	}
	req, resp := account.NewGetMaxAvailSize(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetMaxAvailSizeResponse))
}
