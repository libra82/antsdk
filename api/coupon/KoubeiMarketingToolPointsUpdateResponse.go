package coupon

import (
  "github.com/libra82/antsdk/api"
)

type KoubeiMarketingToolPointsUpdateResponse struct {
  api.AlipayResponse
  PointLogNo string `json:"point_log_no"` // 集点变更流水号
}
