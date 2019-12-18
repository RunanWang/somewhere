package msg

import "github.com/somewhere/model"

type GetBasicReq struct {
}

type GetBasicResp struct {
	model.TBasic
	StdResp
}
