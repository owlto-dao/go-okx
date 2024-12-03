package wallet

import "github.com/owlto-dao/go-okx/rest/api"

func NewGetAllTokenBalancesByAddr(param *GetAllTokenBalancesByAddrParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/wallet/asset/all-token-balances-by-address",
		Method: api.MethodGet,
		Param:  param,
	}, &GetAllTokenBalancesByAddrResponse{}
}

type GetAllTokenBalancesByAddrParam struct {
	Address string `url:"address"`
	Chains  string `url:"chains"`
	Filter  string `url:"filter,omitempty"` // eg. 0-filter risk token, 1-not filter fisk token
}

type GetAllTokenBalancesByAddrResponse struct {
	api.Response
	Data []AllTokenBalances `json:"data"`
}

type AllTokenBalances struct {
	TokenAssets []TokenBalances `json:"tokenAssets"`
}

type TokenBalances struct {
	ChainIndex   string `json:"chainIndex"`
	TokenAddress string `json:"tokenAddress"`
	Address      string `json:"address"`
	Symbol       string `json:"symbol"`
	Balance      string `json:"balance"`
	TokenPrice   string `json:"tokenPrice"` // USD price
	TokenType    string `json:"tokenType"`  // 1-token, 2-铭文
	IsRiskToken  bool   `json:"isRiskToken"`
}
