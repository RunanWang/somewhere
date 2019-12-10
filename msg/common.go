package msg

type StdResp struct {
	ErrorCode int64  `json:"code"`
	ErrorMsg  string `json:"message"`
	RequestID string `json:"request_id"`
}
