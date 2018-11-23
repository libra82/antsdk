package open

import (
  "antsdk/api"
)

type AlipayOpenPublicMessageTotalSendResponse struct {
  api.AlipayResponse
  MessageId string `json:"message_id"`  // 消息ID
}
