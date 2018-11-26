package trade

import "github.com/libra82/antsdk/api"

// AlipayTradeAppPayResponse AlipayTradeAppPayResponse
type AlipayTradeAppPayResponse struct {
	api.AlipayResponse
	OutTradeNO  string `json:"out_trade_no"` // 商户网站唯一订单号
	TradeNo     string `json:"trade_no"`     // 该交易在支付宝系统中的交易流水号。最长64位。
	AppID       string `json:"app_id"`       // 支付宝分配给开发者的应用Id。 	2014072300007148
	TotalAmount string `json:"total_amount"` // 该笔订单的资金总额，单位为RMB-Yuan。取值范围为[0.01,100000000.00]，精确到小数点后两位。 9.00
	SellerID    string `json:"seller_id"`    // 收款支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字 2088111111116894
	Charset     string `json:"charset"`      // 编码格式 utf-8
	Timestamp   string `json:"timestamp"`    // 时间 2016-10-11 17:43:36
}
