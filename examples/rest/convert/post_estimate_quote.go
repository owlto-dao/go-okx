package main

import (
	"log"

	"github.com/owlto-dao/go-okx/examples/rest"
	"github.com/owlto-dao/go-okx/rest/api"
	"github.com/owlto-dao/go-okx/rest/api/convert"
)

func main() {
	param := &convert.PostEstimateQuoteParam{
		BaseCcy:  "BTC",
		QuoteCcy: "USDT",
		Side:     api.SideSell,
		RfqSz:    "10",
		RfqSzCcy: "USDT",
	}
	req, resp := convert.NewPostEstimateQuote(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.PostEstimateQuoteResponse))
}
