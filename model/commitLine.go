package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/config"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/db"
	cerror "gitlab.bj.sensetime.com/SenseGo/cali_server/err"
	"gitlab.bj.sensetime.com/sensego-common/sjwt"
)

type TCommitLine struct {
	Rtsp         string `json:"rtsp" bson:"rtsp"`
	Floor        int    `json:"floor" bson:"floor"`
	FacePosition bool   `json:"facePosition" bson:"facePosition"`
}

type TCommitCalibration struct {
	Calibration [][]float64 `json:"calibration" bson:"calibration"`
}

type TCommitCameraInfo struct {
	DeviceID   string `json:"device_id" bson:"device_id"`
	CameraName string `json:"camera_name" bson:"camera_name"`
	CameraId   string `json:"camera_id" bson:"camera_id"`
}

type TCommitAns struct {
	ProductLine  string      `json:"product_line" bson:"product_line"`
	Ak           string      `json:"ak" bson:"ak"`
	StoreID      string      `json:"store_id" bson:"store_id"`
	Floor        int         `json:"floor" bson:"floor"`
	DeviceID     string      `json:"device_id" bson:"device_id"`
	CameraName   string      `json:"camera_name" bson:"camera_name"`
	CameraId     string      `json:"camera_id" bson:"camera_id"`
	FacePosition bool        `json:"face_position" bson:"face_position"`
	Matrix       [][]float64 `json:"matrix" bson:"matrix"`
	AreaId       int         `json:"area_id" bson:"area_id"`
	LocateMethod string      `json:"locate_method" bson:"locate_method"`
}

type TCameraResp struct {
	ErrorCode int             `json:"error_code"`
	ErrorMsg  string          `json:"error_msg"`
	RequestId string          `json:"request_id"`
	List      []TCameraDetail `json:"list"`
}

type TCameraDetail struct {
	DeviceID    string `json:"device_id"`
	CameraId    string `json:"camera_id"`
	CameraName  string `json:"camera_name"`
	Rtsp        string `json:"rtsp_addr"`
	Ip          string `json:"ip"`
	Offline     bool   `json:"offline"`
	Status_code string `json:"status_code"`
	Timestamp   string `json:"timestamp"`
}

type TCommitResp struct {
	ErrorCode int64  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (t *TCommitLine) CommitGetCalibration(StoreID string) (*TCommitCalibration, error) {
	var ret TCommitCalibration
	col := db.Db.C("cali")
	err := col.Find(bson.M{"store_id": StoreID, "rtsp": t.Rtsp}).One(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (t *TCommitLine) CommitGetCamera(StoreID string, ak string, token string) (*TCommitCameraInfo, error) {
	var ret TCommitCameraInfo
	url := fmt.Sprint(config.Config.ServiceConfig.CameraServer, "?ak=", ak, "&store_id=", StoreID)
	client := &http.Client{}
	req, err := http.NewRequest(`GET`, url, nil)
	if err != nil {
		return &ret, err
	}
	req.Header.Add("token", token)
	resp, err := client.Do(req)
	if err != nil {
		return &ret, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &ret, err
	}

	result := &TCameraResp{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return &ret, err
	}
	if result.ErrorCode != 0 {
		return nil, cerror.ErrMap[int32(result.ErrorCode)]
	}

	found := false
	for _, v := range result.List {
		if v.Rtsp == t.Rtsp {
			found = true
			ret.CameraId = v.CameraId
			ret.CameraName = v.CameraName
			ret.DeviceID = v.DeviceID
		}
	}

	if !found {
		return &ret, cerror.ErrNotFound
	}
	return &ret, nil
}

func CommitGetAk(token string) (string, error) {
	meta, err := sjwt.Decode(token)
	if err != nil {
		return "", err
	}
	return meta.Ak, nil
}

func (t *TCommitLine) CommitFormAns(ak string, StoreID string, cal *TCommitCalibration, camera *TCommitCameraInfo) TCommitAns {
	var tempAns TCommitAns

	tempAns.ProductLine = "2"
	tempAns.StoreID = StoreID
	tempAns.Floor = t.Floor
	tempAns.FacePosition = t.FacePosition
	tempAns.Ak = ak
	tempAns.LocateMethod = "default"
	tempAns.Matrix = cal.Calibration
	tempAns.CameraId = camera.CameraId
	tempAns.CameraName = camera.CameraName
	tempAns.DeviceID = camera.DeviceID

	return tempAns
}

func (t *TCommitLine) CommitChangeStatus(StoreID string) error {
	col := db.Db.C("cali")

	err := col.Update(bson.M{"store_id": StoreID, "rtsp": t.Rtsp}, bson.M{"$set": bson.M{"status": 1}})
	if err != nil {
		return err
	}
	return nil
}

func CommitToConvertServer(toCommit TCommitAns) error {
	jsonStr, err := json.Marshal(toCommit)
	if err != nil {
		return err
	}
	resp, err := http.Post(config.Config.ServiceConfig.ConvertServer, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	result := &TCommitResp{}
	err = json.Unmarshal(b, result)
	if err != nil {
		fmt.Println(string(b))
		return err
	}
	fmt.Println("post get result:", result)
	if result.ErrorCode != 0 {
		return cerror.ErrMap[int32(result.ErrorCode)]
	}
	return nil
}
