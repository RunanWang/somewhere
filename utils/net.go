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

type StdResp struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func GetItemScoreFromUserID(UserID string) (ScoreResp, error) {
	var reqCont ScoreReq
	var respCont ScoreResp
	reqCont.UserID = UserID
	jsonStr, err := json.Marshal(reqCont)
	if err != nil {
		return respCont, err
	}
	url := fmt.Sprint(config.Config.AlgoConfig.Address, "/test")
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

func TrainModel() error {
	var respCont StdResp
	var jsonStr []byte
	url := fmt.Sprint(config.Config.AlgoConfig.Address, "/train")
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(b)
		return err
	}
	err = json.Unmarshal(b, &respCont)
	if err != nil {
		return err
	}
	return nil
}
