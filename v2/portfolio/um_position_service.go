package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeUmLeverageService change user's initial leverage of specific symbol market
type ChangeUmLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeUmLeverageService) Symbol(symbol string) *ChangeUmLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeUmLeverageService) Leverage(leverage int) *ChangeUmLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeUmLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *UmSymbolLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UmSymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UmSymbolLeverage define leverage info of symbol
type UmSymbolLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}
