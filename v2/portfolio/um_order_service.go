package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateUmOrderService create order
type CreateUmOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	priceProtect     *bool
	newOrderRespType NewOrderRespType
	closePosition    *bool
}

// Symbol set symbol
func (s *CreateUmOrderService) Symbol(symbol string) *CreateUmOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateUmOrderService) Side(side SideType) *CreateUmOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateUmOrderService) PositionSide(positionSide PositionSideType) *CreateUmOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateUmOrderService) Type(orderType OrderType) *CreateUmOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateUmOrderService) TimeInForce(timeInForce TimeInForceType) *CreateUmOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateUmOrderService) Quantity(quantity string) *CreateUmOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateUmOrderService) ReduceOnly(reduceOnly bool) *CreateUmOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateUmOrderService) Price(price string) *CreateUmOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateUmOrderService) NewClientOrderID(newClientOrderID string) *CreateUmOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateUmOrderService) StopPrice(stopPrice string) *CreateUmOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *CreateUmOrderService) WorkingType(workingType WorkingType) *CreateUmOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateUmOrderService) ActivationPrice(activationPrice string) *CreateUmOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateUmOrderService) CallbackRate(callbackRate string) *CreateUmOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateUmOrderService) PriceProtect(priceProtect bool) *CreateUmOrderService {
	s.priceProtect = &priceProtect
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateUmOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateUmOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *CreateUmOrderService) ClosePosition(closePosition bool) *CreateUmOrderService {
	s.closePosition = &closePosition
	return s
}

func (s *CreateUmOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateUmOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateUmOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/um/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateUmOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUmOrderResponse define create order response
type CreateUmOrderResponse struct {
	Symbol            string           `json:"symbol"`
	OrderID           int64            `json:"orderId"`
	ClientOrderID     string           `json:"clientOrderId"`
	Price             string           `json:"price"`
	OrigQuantity      string           `json:"origQty"`
	ExecutedQuantity  string           `json:"executedQty"`
	CumQuote          string           `json:"cumQuote"`
	ReduceOnly        bool             `json:"reduceOnly"`
	Status            OrderStatusType  `json:"status"`
	StopPrice         string           `json:"stopPrice"`
	TimeInForce       TimeInForceType  `json:"timeInForce"`
	Type              OrderType        `json:"type"`
	Side              SideType         `json:"side"`
	UpdateTime        int64            `json:"updateTime"`
	WorkingType       WorkingType      `json:"workingType"`
	ActivatePrice     string           `json:"activatePrice"`
	PriceRate         string           `json:"priceRate"`
	AvgPrice          string           `json:"avgPrice"`
	PositionSide      PositionSideType `json:"positionSide"`
	ClosePosition     bool             `json:"closePosition"`
	PriceProtect      bool             `json:"priceProtect"`
	RateLimitOrder10s string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string           `json:"rateLimitOrder1m,omitempty"`
}

// GetUmOrderService get an order
type GetUmOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetUmOrderService) Symbol(symbol string) *GetUmOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetUmOrderService) OrderID(orderID int64) *GetUmOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetUmOrderService) OrigClientOrderID(origClientOrderID string) *GetUmOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetUmOrderService) Do(ctx context.Context, opts ...RequestOption) (res *UmOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UmOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UmOrder define order info
type UmOrder struct {
	Symbol           string           `json:"symbol"`
	OrderID          int64            `json:"orderId"`
	ClientOrderID    string           `json:"clientOrderId"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	OrigQuantity     string           `json:"origQty"`
	ExecutedQuantity string           `json:"executedQty"`
	CumQuantity      string           `json:"cumQty"`
	CumQuote         string           `json:"cumQuote"`
	Status           OrderStatusType  `json:"status"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	Side             SideType         `json:"side"`
	StopPrice        string           `json:"stopPrice"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	AvgPrice         string           `json:"avgPrice"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
	ClosePosition    bool             `json:"closePosition"`
}

// CancelUmOrderService cancel an order
type CancelUmOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *CancelUmOrderService) Symbol(symbol string) *CancelUmOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelUmOrderService) OrderID(orderID int64) *CancelUmOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelUmOrderService) OrigClientOrderID(origClientOrderID string) *CancelUmOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *CancelUmOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelUmOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelUmOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelUmOrderResponse define response of canceling order
type CancelUmOrderResponse struct {
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      string           `json:"cumQty"`
	CumQuote         string           `json:"cumQuote"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     string           `json:"origQty"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	Symbol           string           `json:"symbol"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
}

// CancelAllOpenUmOrdersService cancel all open orders
type CancelAllOpenUmOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelAllOpenUmOrdersService) Symbol(symbol string) *CancelAllOpenUmOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelAllOpenUmOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
