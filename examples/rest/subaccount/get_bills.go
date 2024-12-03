package main

import (
	"log"

	"github.com/owlto-dao/go-okx/examples/rest"
	"github.com/owlto-dao/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetBillsParam{}
	req, resp := subaccount.NewGetBills(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.GetBillsResponse))
}
