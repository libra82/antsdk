package trade

import (
  "antsdk/api"
)

type AlipayTradeOrderSettleResponse struct {
  api.AlipayResponse
  TradeNo string `json:"trade_no"` //支付宝交易号
}
