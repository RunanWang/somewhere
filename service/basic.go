package service

import (
	"github.com/somewhere/model"
)

func BasicInit() error {
	return model.BasicInit()
}

func GetBasic() model.TBasic {
	return model.Basic
}
