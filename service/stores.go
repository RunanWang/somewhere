package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddStore(c *gin.Context, addStoreReq *msg.AddStoresReq) (string, error) {
	storeModel := &model.TStores{
		ID:        bson.NewObjectId(),
		Name:      addStoreReq.Name,
		NickName:  addStoreReq.StoreName,
		Level:     addStoreReq.StoreLevel,
		City:      addStoreReq.StoreCity,
		Timestamp: time.Now().Unix(),
	}
	logger := c.MustGet("logger").(*log.Entry)
	err := storeModel.AddStore()
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"add_item_error": err,
		})
		c.Set("logger", logger)
		return storeModel.ID.Hex(), err
	}
	AuthModel := &model.TAuth{
		ID:       bson.NewObjectId(),
		Username: addStoreReq.Name,
		Password: "111111",
		Role:     "shop",
	}
	err = AuthModel.AddAuth()
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"add_user_auth_error": err,
		})
		c.Set("logger", logger)
		return storeModel.ID.Hex(), err
	}
	return storeModel.ID.Hex(), err
}

func GetStores(c *gin.Context, getStoresReq *msg.GetStoresReq) ([]model.TStores, error) {
	return model.GetAllStores()
}

func GetStoresByPage(c *gin.Context, getStoresByPageReq *msg.GetStoresByPageReq) ([]model.TStores, error) {
	return model.GetStoresByPage(getStoresByPageReq.PageNum, getStoresByPageReq.PageSize)
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
	logger := c.MustGet("logger").(*log.Entry)
	storeModel := &model.TStores{
		ID: bson.ObjectIdHex(delStoreReq.StoreID),
	}
	AuthModel := &model.TAuth{
		Username: delStoreReq.Name,
	}
	err := AuthModel.DeleteAuthByName()
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"del_user_error": err,
		})
		c.Set("logger", logger)
		return 1, err
	}

	return 1, storeModel.DeleteStore()
}
