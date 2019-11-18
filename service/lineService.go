package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/model"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/msg"
)

type TCommitLine struct {
	Rtsp         string `json:"rtsp" bson:"rtsp"`
	Floor        int    `json:"floor" bson:"floor"`
	FacePosition bool   `json:"facePosition" bson:"facePosition"`
}

func GetLine(c *gin.Context, req *msg.GetLineReq, resp *msg.GetLineResp) error {

	logger := c.MustGet("logger").(*log.Entry)

	ret, err := model.GetStoreCalibration(req.StoreID)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"get_cali_error": err.Error(),
		})
		c.Set("logger", logger)
		return err
	}
	resp.List = ret

	return nil
}

func CommitLine(c *gin.Context, req *msg.CommitLineReq, resp *msg.CommitLineResp) error {
	var commitAns []model.TCommitAns
	logger := c.MustGet("logger").(*log.Entry)
	// 1-先获得组合数据，逐条发给converter
	StoreID := req.StoreID
	token := c.Request.Header.Get("token")
	ak, err := model.CommitGetAk(token)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"commit_get_ak": err.Error(),
		})
		c.Set("logger", logger)
		return err
	}
	successCount := 0

	for _, v := range req.List {
		// 获取坐标数组
		cal, err := v.CommitGetCalibration(StoreID)
		if err != nil {
			logger = logger.WithFields(log.Fields{
				"commit_get_calibration": err.Error(),
			})
			c.Set("logger", logger)
			continue
		}

		// 获取摄像头信息
		camera, err := v.CommitGetCamera(StoreID, ak, token)
		if err != nil {
			logger = logger.WithFields(log.Fields{
				"commit_get_camera": err.Error(),
			})
			c.Set("logger", logger)
			continue
		}

		// 形成给converter的报文信息
		tempAns := v.CommitFormAns(ak, StoreID, cal, camera)

		// 发送给converter
		err = model.CommitToConvertServer(tempAns)
		if err != nil {
			logger = logger.WithFields(log.Fields{
				"commit_to_converter": err.Error(),
			})
			c.Set("logger", logger)
			continue
		}
		commitAns = append(commitAns, tempAns)

		//发送成功，改状态并计数
		err = v.CommitChangeStatus(StoreID)
		if err != nil {
			logger = logger.WithFields(log.Fields{
				"commit_change_status": err.Error(),
			})
			c.Set("logger", logger)
			continue
		}
		successCount++
	}

	logger = logger.WithFields(log.Fields{
		"success_count": successCount,
		"commit_items":  commitAns,
	})
	c.Set("logger", logger)

	//3-最后返回成功个数
	resp.Count = successCount
	return nil
}
