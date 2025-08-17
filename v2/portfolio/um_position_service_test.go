package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umPositionServiceTestSuite struct {
	baseTestSuite
}

func TestUmPositionService(t *testing.T) {
	suite.Run(t, new(umPositionServiceTestSuite))
}

func (s *umPositionServiceTestSuite) TestChangeLeverage() {
	data := []byte(`{
		"leverage": 21,
		"maxNotionalValue": "1000000",
		"symbol": "BTCUSDT"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	leverage := 21
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":   symbol,
			"leverage": leverage,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewChangeUmLeverageService().Symbol(symbol).Leverage(leverage).Do(newContext())
	s.r().NoError(err)
	e := &UmSymbolLeverage{
		Symbol:           symbol,
		Leverage:         leverage,
		MaxNotionalValue: "1000000",
	}
	s.r().Equal(e.Symbol, res.Symbol, "Symbol")
	s.r().Equal(e.Leverage, res.Leverage, "Leverage")
	s.r().Equal(e.MaxNotionalValue, res.MaxNotionalValue, "MaxNotionalValue")
}
