package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUmAccountService get account info
type GetUmAccountService struct {
	c *Client
}

// Do send request
func (s *GetUmAccountService) Do(ctx context.Context, opts ...RequestOption) (res *UmAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UmAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UmAccount define account info
type UmAccount struct {
	Assets                      []*UmAccountAsset    `json:"assets"`
	FeeTier                     int                  `json:"feeTier"`
	CanTrade                    bool                 `json:"canTrade"`
	CanDeposit                  bool                 `json:"canDeposit"`
	CanWithdraw                 bool                 `json:"canWithdraw"`
	UpdateTime                  int64                `json:"updateTime"`
	MultiAssetsMargin           bool                 `json:"multiAssetsMargin"`
	TotalInitialMargin          string               `json:"totalInitialMargin"`
	TotalMaintMargin            string               `json:"totalMaintMargin"`
	TotalWalletBalance          string               `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string               `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string               `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string               `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string               `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string               `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string               `json:"totalCrossUnPnl"`
	AvailableBalance            string               `json:"availableBalance"`
	MaxWithdrawAmount           string               `json:"maxWithdrawAmount"`
	Positions                   []*UmAccountPosition `json:"positions"`
}

// UmAccountAsset define account asset
type UmAccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

// UmAccountPosition define account position
type UmAccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               string           `json:"leverage"`
	InitialMargin          string           `json:"initialMargin"`
	MaintMargin            string           `json:"maintMargin"`
	OpenOrderInitialMargin string           `json:"openOrderInitialMargin"`
	PositionInitialMargin  string           `json:"positionInitialMargin"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       string           `json:"unrealizedProfit"`
	EntryPrice             string           `json:"entryPrice"`
	MaxNotional            string           `json:"maxNotional"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            string           `json:"positionAmt"`
	Notional               string           `json:"notional"`
	BidNotional            string           `json:"bidNotional"`
	AskNotional            string           `json:"askNotional"`
	UpdateTime             int64            `json:"updateTime"`
}
