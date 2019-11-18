package msg

import jwt "github.com/dgrijalva/jwt-go"

type TGenerateKeyReq struct {
	StoreID string `form:"store_id" binding:"required"`
}

type TGenerateResp struct {
	StdResp
	Token string `json:"token"`
}

type TTokenMeta struct {
	StoreID string `json:"store_id"`
	Url     string `json:"url"`
	jwt.StandardClaims
}
