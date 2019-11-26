package service

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddStore(c *gin.Context, addStoreReq *msg.AddStoresReq) (string, error) {
	storeModel := &model.TStores{
		ID:    bson.NewObjectId(),
		Name:  addStoreReq.StoreName,
		Level: addStoreReq.StoreLevel,
		City:  addStoreReq.StoreCity,
	}
	logger := c.MustGet("logger").(*log.Entry)
	err := storeModel.AddStore()
	logger = logger.WithFields(log.Fields{
		"add_item_error": err,
	})
	c.Set("logger", logger)
	return storeModel.ID.Hex(), err
}

func GetStores(c *gin.Context, getStoresReq *msg.GetStoresReq) ([]model.TStores, error) {
	return model.GetAllStores()
}

func UpdateStore(c *gin.Context, updateStoresReq *msg.UpdateStoresReq) (int, error) {
	storeModel := &model.TStores{
		ID:    bson.ObjectIdHex(updateStoresReq.StoreID),
		Name:  updateStoresReq.StoreName,
		Level: updateStoresReq.StoreLevel,
		City:  updateStoresReq.StoreCity,
	}

	return 1, storeModel.UpdateStore()
}

func DeleteStore(c *gin.Context, delStoreReq *msg.DeleteStoresReq) (int, error) {
	storeModel := &model.TStores{
		ID: bson.ObjectIdHex(delStoreReq.StoreID),
	}

	return 1, storeModel.DeleteStore()
}
