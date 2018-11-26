package car

import (
  "github.com/libra82/antsdk/api"
)

type AlipayEcoMycarParkingVehicleQueryResponse struct {
  api.AlipayResponse
  CarNumber string `json:"car_number"` // 车牌信息（utf-8编码）
}
