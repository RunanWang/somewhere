package msg

type AddAuthReq struct {
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddAuthResp struct {
	StdResp
}
