package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "gitlab.bj.sensetime.com/SenseGo/cali_server/err"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/model"
)

const (
	FILE_MAX_SIZE = 1024 * 1024
)

type TCaliUnit struct {
	CameraCalibration [][]float64 `json:"camera_calibration"`
	RtspAddress       string      `json:"rtsp_address"`
}

func AddFile(c *gin.Context, logger *log.Entry) error {

	var (
		Content []*TCaliUnit
	)

	storeID, b := c.GetPostForm("store_id")
	if !b {
		return cerror.ErrInvalidParam
	}

	multiFile, err := c.FormFile("file")
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"multifile_error": err.Error(),
		})
		c.Set("logger", logger)
		return err
	}

	file, err := multiFile.Open()
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"oepnfile_error": err.Error(),
		})
		c.Set("logger", logger)
		return err
	}
	defer file.Close()

	// 小于1m 过大的文件不受理
	buf := make([]byte, FILE_MAX_SIZE)

	n, err := file.Read(buf)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"readfile_error": err.Error(),
		})
		c.Set("logger", logger)
		return err
	}
	if n > FILE_MAX_SIZE {
		return cerror.ErrInvalidParam
	}

	err = json.Unmarshal(buf[:n], &Content)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"unmarshal_error": err.Error(),
		})
		c.Set("logger", logger)
		return err
	}
	for _, v := range Content {
		t := new(model.TCalibration)
		t.Status = 0
		t.StoreID = storeID
		t.Rtsp = v.RtspAddress
		t.Calibration = v.CameraCalibration
		err = t.AddCalibration()
		fmt.Println("err", err, t.StoreID, t.Rtsp)
		break
	}

	return nil
}
