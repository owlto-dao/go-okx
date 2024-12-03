package main

import (
	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api/wallet"
	"log"
)

func main() {
	req, resp := wallet.NewGetSupportedChains()
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*wallet.GetSupportedChainsResponse))
}
