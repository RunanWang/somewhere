package msg

type StdResp struct {
	ErrorCode int64  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	RequestID string `json:"request_id"`
}
