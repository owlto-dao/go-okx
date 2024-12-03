package wallet

import (
	"github.com/owlto-dao/go-okx/rest/api"
)

func NewGetTotalValueByAddr(param *GetTotalValueByAddrParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/wallet/asset/total-value-by-address",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTotalValueByAddrResponse{}
}

type GetTotalValueByAddrParam struct {
	Address          string `url:"address"`
	Chains           string `url:"chains"`
	AssetType        string `url:"assetType,omitempty"`
	ExcludeRiskToken bool   `url:"excludeRiskToken,omitempty"`
}

type GetTotalValueByAddrResponse struct {
	api.Response
	Data []TotalValue `json:"data"`
}

type TotalValue struct {
	TotalValue string `json:"totalValue"`
}
