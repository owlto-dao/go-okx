package wallet

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetSupportedChains() (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/wallet/chain/supported-chains",
		Method: api.MethodGet,
	}, &GetSupportedChainsResponse{}
}

type GetSupportedChainsResponse struct {
	api.Response
	Data []Chain `json:"data"`
}

type Chain struct {
	Name       string `json:"name"`
	LogoUrl    string `json:"logoUrl"`
	ShortName  string `json:"shortName"`
	ChainIndex string `json:"chainIndex"`
}
