package service

import (
	"github.com/gin-gonic/gin"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddRecord(c *gin.Context, addRecordReq *msg.AddRecordReq) (int, error) {
	RecordModel := &model.TRecord{
		ProID:  addRecordReq.ProID,
		Status: addRecordReq.Status,
		UserID: addRecordReq.UserID,
	}

	return RecordModel.AddRecord()
}

func GetRecords(c *gin.Context, getRecordsReq *msg.GetRecordsReq) ([]*model.TRecord, error) {
	//if getRecordsReq.RecordID <= 0 {
	return model.GetAllRecords()
	//} else {
	// RecordsModel := &model.TRecord{
	// 	RecID: getRecordsReq.RecordID,
	// }
	// return model.GetAllRecords()
	//}
}
