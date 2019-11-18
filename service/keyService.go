package service

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/config"
	cerror "gitlab.bj.sensetime.com/SenseGo/cali_server/err"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/msg"
)

const (
	KEY      = "825867ae01a71eb945588a4886df6a26"
	EXE_PATH = "/sensego/calibration/exefile"
)

func GenerateKey(c *gin.Context, generateKeyReq *msg.TGenerateKeyReq, generateKeyResp *msg.TGenerateResp) error {

	expDuration := time.Duration(config.Config.ServiceConfig.ExpireTokenHour)
	Exp := time.Now().Add(expDuration * time.Hour).Unix()

	meta := msg.TTokenMeta{
		generateKeyReq.StoreID,
		joinExePath(c),

		jwt.StandardClaims{
			ExpiresAt: Exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, meta)
	key := []byte(KEY)

	str, err := token.SignedString(key)
	if err != nil {
		return err
	}
	generateKeyResp.Token = str
	return nil
}

func joinExePath(c *gin.Context) string {
	return fmt.Sprintf("%s", c.Request.Host+EXE_PATH)
}

func CheckExeKey(key string) error {
	tokenType, err := jwt.ParseWithClaims(key, &msg.TTokenMeta{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})

	if err != nil {
		return err
	}

	if _, ok := tokenType.Claims.(*msg.TTokenMeta); ok && tokenType.Valid {
		return nil
	}
	return cerror.ErrInvalidParam
}
