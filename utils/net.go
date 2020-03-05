package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/somewhere/config"
)

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

func GetItemScoreFromUserID(UserID string) (ScoreResp, error) {
	var reqCont ScoreReq
	var respCont ScoreResp
	reqCont.UserID = UserID
	jsonStr, err := json.Marshal(reqCont)
	if err != nil {
		return respCont, err
	}
	url := config.Config.AlgoConfig.Address
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return respCont, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return respCont, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(b)
		return respCont, err
	}
	err = json.Unmarshal(b, &respCont)
	if err != nil {
		return respCont, err
	}
	return respCont, nil
}
