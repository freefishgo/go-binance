package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateCMOrderService https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/trade/New-CM-Order 币本位合约下单
type CreateCMOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	reduceOnly              *bool
	price                   *string
	newClientOrderID        *string
	stopPrice               *string
	closePosition           *bool
	activationPrice         *string
	callbackRate            *string
	workingType             *WorkingType
	priceProtect            *bool
	newOrderRespType        NewOrderRespType
	selfTradePreventionMode SelfTradePreventionMode
}

// Symbol set symbol
func (s *CreateCMOrderService) Symbol(symbol string) *CreateCMOrderService {
	s.symbol = symbol
	return s
}

// GetSymbol get symbol
func (s *CreateCMOrderService) GetSymbol() string {
	return s.symbol
}

// Side set side
func (s *CreateCMOrderService) Side(side SideType) *CreateCMOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateCMOrderService) PositionSide(positionSide PositionSideType) *CreateCMOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateCMOrderService) Type(orderType OrderType) *CreateCMOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateCMOrderService) TimeInForce(timeInForce TimeInForceType) *CreateCMOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateCMOrderService) Quantity(quantity string) *CreateCMOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateCMOrderService) ReduceOnly(reduceOnly bool) *CreateCMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateCMOrderService) Price(price string) *CreateCMOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateCMOrderService) NewClientOrderID(newClientOrderID string) *CreateCMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// GetClientOrderID get newClientOrderID
func (s *CreateCMOrderService) GetClientOrderID() string {
	if s.newClientOrderID != nil {
		return *s.newClientOrderID
	}
	return ""
}

// StopPrice set stopPrice
func (s *CreateCMOrderService) StopPrice(stopPrice string) *CreateCMOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *CreateCMOrderService) WorkingType(workingType WorkingType) *CreateCMOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateCMOrderService) ActivationPrice(activationPrice string) *CreateCMOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateCMOrderService) CallbackRate(callbackRate string) *CreateCMOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateCMOrderService) PriceProtect(priceProtect bool) *CreateCMOrderService {
	s.priceProtect = &priceProtect
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateCMOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateCMOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// SelfTradePreventionMode set SelfTradePreventionMode
func (s *CreateCMOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateCMOrderService {
	s.selfTradePreventionMode = selfTradePreventionMode
	return s
}

// ClosePosition set closePosition
func (s *CreateCMOrderService) ClosePosition(closePosition bool) *CreateCMOrderService {
	s.closePosition = &closePosition
	return s
}

func (s *CreateCMOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
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
	if s.selfTradePreventionMode != "" {
		m["selfTradePreventionMode"] = s.selfTradePreventionMode
	}

	r.setFormParams(m)
	data, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateCMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/papi/v1/cm/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//type CreateBatchOrdersOrderService struct {
//	c *Client
//}
//
//// Do send request
//func (c *CreateBatchOrdersOrderService) Do(ctx context.Context, orderList []*CreateCMOrderService, opts ...RequestOption) (res []*CreateOrderResponse, err error) {
//	data, err := c.createOrder(ctx, "/dapi/v1/batchOrders", orderList, opts...)
//	if err != nil {
//		return nil, err
//	}
//	err = json.Unmarshal(data, &res)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//
//func (c *CreateBatchOrdersOrderService) createOrder(ctx context.Context, endpoint string, orderList []*CreateCMOrderService, opts ...RequestOption) (data []byte, err error) {
//	r := &request{
//		method:   http.MethodPost,
//		endpoint: endpoint,
//		secType:  secTypeSigned,
//	}
//	var mList []params
//	for _, s := range orderList {
//		m := params{
//			"symbol":           s.symbol,
//			"side":             s.side,
//			"type":             s.orderType,
//			"quantity":         s.quantity,
//			"newOrderRespType": s.newOrderRespType,
//		}
//		if s.positionSide != nil {
//			m["positionSide"] = *s.positionSide
//		}
//		if s.timeInForce != nil {
//			m["timeInForce"] = *s.timeInForce
//		}
//		if s.reduceOnly != nil {
//			m["reduceOnly"] = *s.reduceOnly
//		}
//		if s.price != nil {
//			m["price"] = *s.price
//		}
//		if s.newClientOrderID != nil {
//			m["newClientOrderId"] = *s.newClientOrderID
//		}
//		if s.stopPrice != nil {
//			m["stopPrice"] = *s.stopPrice
//		}
//		if s.workingType != nil {
//			m["workingType"] = *s.workingType
//		}
//		if s.priceProtect != nil {
//			m["priceProtect"] = *s.priceProtect
//		}
//		if s.activationPrice != nil {
//			m["activationPrice"] = *s.activationPrice
//		}
//		if s.callbackRate != nil {
//			m["callbackRate"] = *s.callbackRate
//		}
//		if s.closePosition != nil {
//			m["closePosition"] = *s.closePosition
//		}
//		mList = append(mList, m)
//	}
//	b, _ := json.Marshal(mList)
//	m := params{
//		"batchOrders": string(b),
//	}
//	r.setFormParams(m)
//	data, err = c.c.callAPI(ctx, r, opts...)
//	if err != nil {
//		return []byte{}, err
//	}
//	return data, nil
//}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      string           `json:"cumQty"`
	CumBase          string           `json:"cumBase"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	AvgPrice         string           `json:"avgPrice"`
	OrigQuantity     string           `json:"origQty"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	ClosePosition    bool             `json:"closePosition"`
	Symbol           string           `json:"symbol"`
	Pair             string           `json:"pair"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	OrigType         OrderType        `json:"origType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	PriceProtect     bool             `json:"priceProtect"`
	Code             int              `json:"code"`
	Msg              string           `json:"msg"`
}

// GetCMOrderService get an order
type GetCMOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetCMOrderService) Symbol(symbol string) *GetCMOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetCMOrderService) OrderID(orderID int64) *GetCMOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetCMOrderService) OrigClientOrderID(origClientOrderID string) *GetCMOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetCMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/order",
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
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Order define order info
type Order struct {
	AvgPrice         string           `json:"avgPrice"`
	ClientOrderID    string           `json:"clientOrderId"`
	CumBase          string           `json:"cumBase"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     string           `json:"origQty"`
	OrigType         OrderType        `json:"origType"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	ClosePosition    bool             `json:"closePosition"`
	Symbol           string           `json:"symbol"`
	Pair             string           `json:"pair"`
	Time             int64            `json:"time"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	PriceProtect     bool             `json:"priceProtect"`
}

// CancelAllCMOrderService https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/trade/Cancel-All-CM-Open-Orders
type CancelAllCMOrderService struct {
	c      *Client
	symbol string
}

type CancelAllCMOrder struct {
}

func (c *CancelAllCMOrderService) Symbol(symbol string) *CancelAllCMOrderService {
	c.symbol = symbol
	return c
}

func (c *CancelAllCMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelAllCMOrder, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/cm/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", c.symbol)
	data, _, err := c.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelAllCMOrder)
	_ = json.Unmarshal(data, res)
	return res, nil
}
