package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddRecord(c *gin.Context, addRecordReq *msg.AddRecordReq) error {
	startTime := time.Now().Unix()
	RecordModel := &model.TRecord{
		ItemID:    addRecordReq.ProID,
		Status:    addRecordReq.Status,
		UserID:    addRecordReq.UserID,
		Timestamp: startTime,
	}
	return RecordModel.AddRecord()
}

func GetRecords(c *gin.Context, getRecordsReq *msg.GetRecordsReq) ([]model.TRecord, error) {
	return model.GetAllRecords()
}
