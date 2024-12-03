package main

import (
	"github.com/owlto-dao/go-okx/examples/rest"
	"github.com/owlto-dao/go-okx/rest/api/wallet"
	"log"
)

func main() {
	req, resp := wallet.NewGetSupportedChains()
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*wallet.GetSupportedChainsResponse))
}
