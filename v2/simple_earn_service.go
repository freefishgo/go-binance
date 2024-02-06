package binance

import (
	"context"
	"net/http"
)

type GetSimpleEarnFlexibleListService struct {
	c       *Client
	asset   *string
	current *string //当前查询页。 开始值 1，默认:1
	size    *string //每页条数。 开始值 10，默认:10
	//recvWindow int64
	//timestamp  int64
}

type SimpleEarnFlexibleList struct {
	Rows  []*SimpleEarnFlexible `json:"rows"`
	Total int                   `json:"total"`
}

type SimpleEarnFlexible struct {
	Asset                      string            `json:"asset"`
	LatestAnnualPercentageRate string            `json:"latestAnnualPercentageRate"`
	TierAnnualPercentageRate   map[string]string `json:"tierAnnualPercentageRate"`
	AirDropPercentageRate      string            `json:"airDropPercentageRate"`
	CanPurchase                bool              `json:"canPurchase"`
	CanRedeem                  bool              `json:"canRedeem"`
	IsSoldOut                  bool              `json:"isSoldOut"`
	Hot                        bool              `json:"hot"`
	MinPurchaseAmount          string            `json:"minPurchaseAmount"`
	ProductId                  string            `json:"productId"`
	SubscriptionStartTime      int64             `json:"subscriptionStartTime"`
	Status                     string            `json:"status"`
}

func (s *GetSimpleEarnFlexibleListService) Asset(asset string) *GetSimpleEarnFlexibleListService {
	s.asset = &asset
	return s
}

func (s *GetSimpleEarnFlexibleListService) Current(current string) *GetSimpleEarnFlexibleListService {
	s.current = &current
	return s
}

func (s *GetSimpleEarnFlexibleListService) Size(size string) *GetSimpleEarnFlexibleListService {
	s.size = &size
	return s
}

func (s *GetSimpleEarnFlexibleListService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexibleList, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/flexible/list",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SimpleEarnFlexibleList)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnFlexibleSubscribeService struct {
	c             *Client
	productId     *string
	amount        *string
	autoSubscribe *bool
	sourceAccount *string //NO	SPOT,FUND,ALL, 默认 SPOT
}

type SimpleEarnFlexibleSubscribe struct {
	PurchaseId int  `json:"purchaseId"`
	Success    bool `json:"success"`
}

func (s *SimpleEarnFlexibleSubscribeService) ProductId(productId string) *SimpleEarnFlexibleSubscribeService {
	s.productId = &productId
	return s
}

func (s *SimpleEarnFlexibleSubscribeService) Amount(amount string) *SimpleEarnFlexibleSubscribeService {
	s.amount = &amount
	return s
}

func (s *SimpleEarnFlexibleSubscribeService) AutoSubscribe(autoSubscribe bool) *SimpleEarnFlexibleSubscribeService {
	s.autoSubscribe = &autoSubscribe
	return s
}

func (s *SimpleEarnFlexibleSubscribeService) SourceAccount(sourceAccount string) *SimpleEarnFlexibleSubscribeService {
	s.sourceAccount = &sourceAccount
	return s
}

func (s *SimpleEarnFlexibleSubscribeService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexibleSubscribe, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/simple-earn/flexible/subscribe",
		secType:  secTypeSigned,
	}
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	if s.amount != nil {
		r.setParam("amount", *s.amount)
	}
	if s.autoSubscribe != nil {
		r.setParam("autoSubscribe", *s.autoSubscribe)
	}
	if s.sourceAccount != nil {
		r.setParam("sourceAccount", *s.sourceAccount)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	res = new(SimpleEarnFlexibleSubscribe)
	err = json.Unmarshal(data, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// SimpleEarnFlexibleRedeemService https://binance-docs.github.io/apidocs/spot/cn/#trade-16
type SimpleEarnFlexibleRedeemService struct {
	c           *Client
	productId   *string //	NO	必填
	redeemAll   *bool   //	NO	true 或者 false, 默认 false
	amount      *string //DECIMAL	NO	当redeemAll为false时必填
	destAccount *string //ENUM	NO	SPOT,FUND,ALL, 默认 SPOT
}

type SimpleEarnFlexibleRedeem struct {
	RedeemId int  `json:"redeemId"`
	Success  bool `json:"success"`
}

func (s *SimpleEarnFlexibleRedeemService) ProductId(productId string) *SimpleEarnFlexibleRedeemService {
	s.productId = &productId
	return s
}

func (s *SimpleEarnFlexibleRedeemService) RedeemAll(redeemAll bool) *SimpleEarnFlexibleRedeemService {
	s.redeemAll = &redeemAll
	return s
}

func (s *SimpleEarnFlexibleRedeemService) Amount(amount string) *SimpleEarnFlexibleRedeemService {
	s.amount = &amount
	return s
}

func (s *SimpleEarnFlexibleRedeemService) DestAccount(destAccount string) *SimpleEarnFlexibleRedeemService {
	s.destAccount = &destAccount
	return s
}

func (s *SimpleEarnFlexibleRedeemService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexibleRedeem, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/simple-earn/flexible/redeem",
		secType:  secTypeSigned,
	}
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	if s.redeemAll != nil {
		r.setParam("redeemAll", *s.redeemAll)
	}
	if s.amount != nil {
		r.setParam("amount", *s.amount)
	}
	if s.destAccount != nil {
		r.setParam("destAccount", *s.destAccount)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	res = new(SimpleEarnFlexibleRedeem)
	err = json.Unmarshal(data, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

type GetSimpleEarnFlexiblePositionService struct {
	c         *Client
	asset     *string //STRING	NO
	productId *string //STRING	NO
	current   *int    //LONG	NO	当前查询页。 开始值 1，默认:1
	size      *int    //LONG	NO	默认：10，最大：100
}

type SimpleEarnFlexiblePositionList struct {
	Rows  []*SapiSimpleEarnFlexiblePosition `json:"rows"`
	Total int                               `json:"total"`
}

type SapiSimpleEarnFlexiblePosition struct {
	TotalAmount                    string            `json:"totalAmount"`
	TierAnnualPercentageRate       map[string]string `json:"tierAnnualPercentageRate"` // "0-5BTC": 0.05, "5-10BTC": 0.03
	LatestAnnualPercentageRate     string            `json:"latestAnnualPercentageRate"`
	YesterdayAirdropPercentageRate string            `json:"yesterdayAirdropPercentageRate"`
	Asset                          string            `json:"asset"`
	AirDropAsset                   string            `json:"airDropAsset"`
	CanRedeem                      bool              `json:"canRedeem"`
	CollateralAmount               string            `json:"collateralAmount"`
	ProductId                      string            `json:"productId"`
	YesterdayRealTimeRewards       string            `json:"yesterdayRealTimeRewards"`
	CumulativeBonusRewards         string            `json:"cumulativeBonusRewards"`
	CumulativeRealTimeRewards      string            `json:"cumulativeRealTimeRewards"`
	CumulativeTotalRewards         string            `json:"cumulativeTotalRewards"`
	AutoSubscribe                  bool              `json:"autoSubscribe"`
}

func (s *GetSimpleEarnFlexiblePositionService) Asset(asset string) *GetSimpleEarnFlexiblePositionService {
	s.asset = &asset
	return s
}

func (s *GetSimpleEarnFlexiblePositionService) ProductId(productId string) *GetSimpleEarnFlexiblePositionService {
	s.productId = &productId
	return s
}

func (s *GetSimpleEarnFlexiblePositionService) Current(current int) *GetSimpleEarnFlexiblePositionService {
	s.current = &current
	return s
}

func (s *GetSimpleEarnFlexiblePositionService) Size(size int) *GetSimpleEarnFlexiblePositionService {
	s.size = &size
	return s
}

func (s *GetSimpleEarnFlexiblePositionService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexiblePositionList, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/flexible/position",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	res = new(SimpleEarnFlexiblePositionList)
	err = json.Unmarshal(data, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
