package service

import (
	"time"

	"github.com/globalsign/mgo/bson"

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
	var tempRec model.TRecord

	if getRecordsReq.ProductID != "" {
		tempRec.ItemID = bson.ObjectIdHex(getRecordsReq.ProductID)
		return tempRec.GetRecordsByItemID()
	}
	if getRecordsReq.UserID != "" {
		tempRec.UserID = bson.ObjectIdHex(getRecordsReq.UserID)
		return tempRec.GetRecordsByUserID()
	}
	return model.GetAllRecords()
}
