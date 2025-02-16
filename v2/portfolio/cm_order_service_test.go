package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseOrderTestSuite struct {
	baseTestSuite
}

type orderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestCreateOrder() {
	//data := []byte(`{
	//	"clientOrderId": "testOrder",
	//	"cumQty": "0",
	//	"cumBase": "0",
	//	"executedQty": "0",
	//	"orderId": 22542179,
	//	"avgPrice": "0.0",
	//	"origQty": "10",
	//	"price": "0",
	//	"reduceOnly": false,
	//	"side": "BUY",
	//	"positionSide": "SHORT",
	//	"status": "NEW",
	//	"stopPrice": "9300",
	//	"closePosition": false,
	//	"symbol": "BTCUSD_200925",
	//	"pair": "BTCUSD",
	//	"timeInForce": "GTC",
	//	"type": "TRAILING_STOP_MARKET",
	//	"origType": "TRAILING_STOP_MARKET",
	//	"activatePrice": "9020",
	//	"priceRate": "0.3",
	//	"updateTime": 1566818724722,
	//	"workingType": "CONTRACT_PRICE",
	//	"priceProtect": false
	//}`)
	//s.mockDo(data, nil)
	//defer s.assertDo()
	symbol := "BNBUSD_PERP"
	side := SideTypeBuy
	positionSide := PositionSideTypeLong
	orderType := OrderTypeMarket
	//timeInForce := TimeInForceTypeFOK
	quantity := "1"
	//reduceOnly := false
	//price := "0"
	newClientOrderID := "testOrder"
	//stopPrice := "9300"
	//closePosition := false
	//activationPrice := "9020"
	//callbackRate := "0.3"
	//workingType := WorkingTypeContractPrice
	priceProtect := false
	newOrderResponseType := NewOrderRespTypeRESULT
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":       symbol,
			"side":         side,
			"positionSide": positionSide,
			"type":         orderType,
			//"timeInForce":  timeInForce,
			"quantity": quantity,
			//"reduceOnly":       reduceOnly,
			//"price":            price,
			"newClientOrderId": newClientOrderID,
			//"stopPrice":        stopPrice,
			//"closePosition":    closePosition,
			//"activationPrice":  activationPrice,
			//"callbackRate":     callbackRate,
			//"workingType":      workingType,
			"priceProtect": priceProtect,
			//"newOrderRespType": newOrderResponseType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateCMOrderService().
		Symbol(symbol).Side(side).Type(orderType).
		//TimeInForce(timeInForce).
		Quantity(quantity).
		//ClosePosition(closePosition).ReduceOnly(reduceOnly).Price(price).
		NewClientOrderID(newClientOrderID).
		//StopPrice(stopPrice).WorkingType(workingType).ActivationPrice(activationPrice).CallbackRate(callbackRate).
		PositionSide(positionSide).
		PriceProtect(priceProtect).
		NewOrderResponseType(newOrderResponseType).
		Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		ClientOrderID:    newClientOrderID,
		CumQuantity:      "0",
		CumBase:          "0",
		ExecutedQuantity: "0",
		OrderID:          22542179,
		AvgPrice:         "0.0",
		OrigQuantity:     "10",
		Price:            "0",
		ReduceOnly:       false,
		Side:             SideTypeBuy,
		PositionSide:     positionSide,
		Status:           OrderStatusTypeNew,
		//StopPrice:        stopPrice,
		ClosePosition: false,
		Symbol:        symbol,
		Pair:          "BTCUSD",
		TimeInForce:   TimeInForceTypeGTC,
		Type:          OrderTypeTrailingStopMarket,
		OrigType:      OrderTypeTrailingStopMarket,
		//ActivatePrice:    activationPrice,
		//PriceRate:        callbackRate,
		UpdateTime:   1566818724722,
		PriceProtect: priceProtect,
		WorkingType:  WorkingTypeContractPrice,
	}
	s.assertCreateOrderResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCreateOrderResponseEqual(e, a *CreateOrderResponse) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CumQuantity, a.CumQuantity, "CumQuantity")
	r.Equal(e.CumBase, a.CumBase, "CumBase")
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.OrigType, a.OrigType, "OrigType")
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
