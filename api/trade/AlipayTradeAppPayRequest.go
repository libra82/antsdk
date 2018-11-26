package trade

import (
	"github.com/libra82/antsdk/api"
	"github.com/libra82/antsdk/utils"
)

// AlipayTradeAppPayRequest alipay.trade.app.pay(App支付请求)
// https://docs.open.alipay.com/204/105465/
// 用于交易创建后，用户在一定时间内未进行支付，可调用该接口直接将未付款的交易进行关闭。
type AlipayTradeAppPayRequest struct {
	api.IAlipayRequest

	NotifyUrl  string                             `json:"notify_url"`
	ReturnUrl  string                             `json:"return_url"`
	BizContent AlipayTradeAppPayRequestBizContent `json:"biz_content"`
}

type AlipayTradeAppPayRequestBizContent struct {
	Body               string            `json:"body"`                 // 可选 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	Subject            string            `json:"subject"`              // 必选 商品的标题/交易标题/订单标题/订单关键字等。
	OutTradeNo         string            `json:"out_trade_no"`         // 必选 商户网站唯一订单号
	TotalAmount        string            `json:"total_amount"`         // 必选 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	ProductCode        string            `json:"product_code"`         // 必选 销售产品码，商家和支付宝签约的产品码，为固定值QUICK_MSECURITY_PAY
	TimeoutExpress     string            `json:"timeout_express"`      // 可选 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。 注：若为空，则默认为15d。
	//SellerId            string            `json:"seller_id"`             // 可选 商户门店编号。该参数用于请求参数中以区分各门店，非必传项。
	//GoodsType          string            `json:"goods_type"`           // 可选 商品主类型：0—虚拟类商品，1—实物类商品 注：虚拟类商品不支持使用花呗渠道
	//PassbackParams     string            `json:"passback_params"`      // 可选 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝
	//PromoParams        string            `json:"promo_params"`         // 可选 优惠参数 注：仅与支付宝协商后可用
	//ExtendParams       *ATAPExtendParams `json:"extend_params"`        // 可选 业务扩展参数，详见下面的“业务扩展参数说明”
	//EnablePayChannels  string            `json:"enable_pay_channels"`  // 可选 可用渠道，用户只能在指定渠道范围内支付 当有多个渠道时用“,”分隔 注：与disable_pay_channels互斥
	//DisablePayChannels string            `json:"disable_pay_channels"` // 可选 禁用渠道，用户不可用指定渠道支付 当有多个渠道时用“,”分隔 注：与enable_pay_channels互斥
	//StoreId            string            `json:"store_id"`             // 可选 商户门店编号。该参数用于请求参数中以区分各门店，非必传项。
}
type ATAPExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id"` // 可选 系统商编号，该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	NeedBuyerRealnamed   string `json:"needBuyerRealnamed"`      // 可选 是否发起实名校验 T：发起 F：不发起
	TransMemo            string `json:"TRANS_MEMO"`              // 可选 账务备注 注：该字段显示在离线账单的账务备注中
	HBFQNum              string `json:"hb_fq_num"`               // 可选 花呗分期数（目前仅支持3、6、12）注：使用该参数需要仔细阅读“花呗分期接入文档”
	HBFQSellerPercent    string `json:"hb_fq_seller_percent"`    // 可选 卖家承担收费比例，商家承担手续费传入100，用户承担手续费传入0，仅支持传入100、0两种，其他比例暂不支持 注：使用该参数需要仔细阅读“花呗分期接入文档”
}

func (this *AlipayTradeAppPayRequest) GetApiMethodName() string {
	return "alipay.trade.app.pay"
}

func (this *AlipayTradeAppPayRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayTradeAppPayRequest) GetTerminalType() string {
	return ""
}

func (this *AlipayTradeAppPayRequest) GetTerminalInfo() string {
	return ""
}

func (this *AlipayTradeAppPayRequest) GetNotifyUrl() string {
	return this.NotifyUrl
}

func (this *AlipayTradeAppPayRequest) GetReturnUrl() string {
	return this.ReturnUrl
}

func (this *AlipayTradeAppPayRequest) GetProdCode() string {
	return "QUICK_MSECURITY_PAY"
}

func (this *AlipayTradeAppPayRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayTradeAppPayRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJson(this.BizContent))
	return txtParams
}
