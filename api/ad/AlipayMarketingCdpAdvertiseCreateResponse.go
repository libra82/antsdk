package ad

import (
  "github.com/libra82/antsdk/api"
)

type AlipayMarketingCdpAdvertiseCreateResponse struct {
  api.AlipayResponse
  AdId string `json:"ad_id"`  // 创建广告唯一标识
}
