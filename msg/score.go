package msg

type ScoreDetail struct {
	ItemID string  `json:"item_id"`
	Score  float64 `json:"score"`
}

type ScoreResp struct {
	List []ScoreDetail `json:"msg"`
}

type ScoreReq struct {
	UserID string `json:"user_id"`
}
