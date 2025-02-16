package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCMAccountService get account info
type GetCMAccountService struct {
	c *Client
}

// Do send request
func (s *GetCMAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CMAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CMAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMAccount define account info
type CMAccount struct {
	Assets      []*CMAccountAsset    `json:"assets"`
	CanDeposit  bool                 `json:"canDeposit"`
	CanTrade    bool                 `json:"canTrade"`
	CanWithdraw bool                 `json:"canWithdraw"`
	FeeTier     int                  `json:"feeTier"`
	Positions   []*CMAccountPosition `json:"positions"`
	UpdateTime  int64                `json:"updateType"`
}

// CMAccountAsset define account asset
type CMAccountAsset struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
}

// CMAccountPosition define accoutn position
type CMAccountPosition struct {
	Symbol                 string `json:"symbol"`
	PositionAmt            string `json:"positionAmt"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	Isolated               bool   `json:"isolated"`
	PositionSide           string `json:"positionSide"`
	EntryPrice             string `json:"entryPrice"`
	MaxQty                 string `json:"maxQty"`
}
