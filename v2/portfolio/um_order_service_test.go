package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestUmOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestCreateUmOrder() {
	data := []byte(`{
		"clientOrderId": "testOrder",
		"cumQuote": "0",
		"executedQty": "0",
		"orderId": 22542179,
		"origQty": "10",
		"price": "10000",
		"reduceOnly": false,
		"side": "SELL",
		"status": "NEW",
		"stopPrice": "0",
		"symbol": "BTCUSDT",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"updateTime": 1566818724722,
		"workingType": "CONTRACT_PRICE",
		"activatePrice": "1000",
		"priceRate": "0.1",
		"positionSide": "BOTH",
		"closePosition": false,
		"priceProtect": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	side := SideTypeSell
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	positionSide := PositionSideTypeBoth
	quantity := "10"
	price := "10000"
	newClientOrderID := "testOrder"
	reduceOnly := false
	stopPrice := "0"
	activationPrice := "1000"
	callbackRate := "0.1"
	workingType := WorkingTypeContractPrice
	priceProtect := true
	newOrderResponseType := NewOrderRespTypeRESULT
	closePosition := false
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"timeInForce":      timeInForce,
			"positionSide":     positionSide,
			"quantity":         quantity,
			"reduceOnly":       reduceOnly,
			"price":            price,
			"newClientOrderId": newClientOrderID,
			"stopPrice":        stopPrice,
			"workingType":      workingType,
			"activationPrice":  activationPrice,
			"callbackRate":     callbackRate,
			"priceProtect":     priceProtect,
			"newOrderRespType": newOrderResponseType,
			"closePosition":    closePosition,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateUmOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).ClosePosition(closePosition).
		ReduceOnly(reduceOnly).Price(price).NewClientOrderID(newClientOrderID).
		StopPrice(stopPrice).WorkingType(workingType).ActivationPrice(activationPrice).
		CallbackRate(callbackRate).PositionSide(positionSide).
		PriceProtect(priceProtect).NewOrderResponseType(newOrderResponseType).
		Do(newContext())
	s.r().NoError(err)
	e := &CreateUmOrderResponse{
		ClientOrderID:    newClientOrderID,
		CumQuote:         "0",
		ExecutedQuantity: "0",
		OrderID:          22542179,
		OrigQuantity:     "10",
		PositionSide:     positionSide,
		Price:            "10000",
		ReduceOnly:       false,
		Side:             SideTypeSell,
		Status:           OrderStatusTypeNew,
		StopPrice:        "0",
		Symbol:           symbol,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		UpdateTime:       1566818724722,
		WorkingType:      WorkingTypeContractPrice,
		ActivatePrice:    activationPrice,
		PriceRate:        callbackRate,
		ClosePosition:    false,
		PriceProtect:     priceProtect,
	}
	s.assertCreateUmOrderResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCreateUmOrderResponseEqual(e, a *CreateUmOrderResponse) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CumQuote, a.CumQuote, "CumQuote")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.ActivatePrice, a.ActivatePrice, "ActivatePrice")
	r.Equal(e.PriceRate, a.PriceRate, "PriceRate")
	r.Equal(e.ClosePosition, a.ClosePosition, "ClosePosition")
}

func (s *orderServiceTestSuite) TestGetUmOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"price": "0.1",
		"reduceOnly": false,
		"origQty": "1.0",
		"executedQty": "0.0",
		"cumQuote": "0.0",
		"status": "NEW",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"stopPrice": "0.0",
		"time": 1499827319559,
		"updateTime": 1499827319559,
		"workingType": "CONTRACT_PRICE",
		"activatePrice": "10000",
		"priceRate":"0.1",
		"positionSide": "BOTH",
		"priceProtect": false,
		"closePosition": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	orderID := int64(1)
	origClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"orderId":           orderID,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetUmOrderService().Symbol(symbol).
		OrderID(orderID).OrigClientOrderID(origClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &UmOrder{
		Symbol:           symbol,
		OrderID:          1,
		ClientOrderID:    origClientOrderID,
		Price:            "0.1",
		ReduceOnly:       false,
		OrigQuantity:     "1.0",
		ExecutedQuantity: "0.0",
		CumQuote:         "0.0",
		Status:           OrderStatusTypeNew,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		Side:             SideTypeBuy,
		StopPrice:        "0.0",
		Time:             1499827319559,
		UpdateTime:       1499827319559,
		WorkingType:      WorkingTypeContractPrice,
		ActivatePrice:    "10000",
		PriceRate:        "0.1",
		PositionSide:     "BOTH",
		PriceProtect:     false,
		ClosePosition:    true,
	}
	s.assertUmOrderEqual(e, order)
}

func (s *baseOrderTestSuite) assertUmOrderEqual(e, a *UmOrder) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CumQuantity, a.CumQuantity, "CumQuantity")
	r.Equal(e.CumQuote, a.CumQuote, "CumQuote")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.Time, e.Time, "Time")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.ActivatePrice, a.ActivatePrice, "ActivatePrice")
	r.Equal(e.PriceRate, a.PriceRate, "PriceRate")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
}

func (s *orderServiceTestSuite) TestCancelAllOpenOrders() {
	data := []byte(`{
		"code": "200",
		"msg": "The operation of cancel all open order is done."
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewCancelAllUmOrderService().Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
}
