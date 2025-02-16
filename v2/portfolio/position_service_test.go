package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type positionServiceTestSuite struct {
	baseTestSuite
}

func TestPositionService(t *testing.T) {
	suite.Run(t, new(positionServiceTestSuite))
}

func (s *positionServiceTestSuite) TestChangeLeverage() {
	//data := []byte(`{
	//	"leverage": 21,
	//	"maxQty": "1000",
	//	"symbol": "BTCUSD_200925"
	//}`)
	//s.mockDo(data, nil)
	//defer s.assertDo()
	symbol := "BNBUSD_PERP"
	leverage := 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":   symbol,
			"leverage": leverage,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewChangeLeverageService().Symbol(symbol).Leverage(leverage).Do(newContext())
	s.r().NoError(err)
	e := &SymbolLeverage{
		Symbol:      symbol,
		Leverage:    leverage,
		MaxQuantity: "1000",
	}
	s.r().Equal(e.Symbol, res.Symbol, "Symbol")
	s.r().Equal(e.Leverage, res.Leverage, "Leverage")
	s.r().Equal(e.MaxQuantity, res.MaxQuantity, "MaxQuantity")
}
