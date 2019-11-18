package cerror

import "fmt"

type Error struct {
	Code int    `json:"error_code"`
	Msg  string `json:"error_msg"`
}

// RPC error variables should consistent with defaultErrMap defined bellow
var (
	ErrMethodNotSupport          = NewError(1001, "method not support")
	ErrInvalidParam              = NewError(1002, "parameter is invalid")
	ErrPermissionDenied          = NewError(1003, "permission denied")
	ErrInvalidToken              = NewError(1004, "invalid token")
	ErrExpired                   = NewError(1005, "expired")
	ErrSignError                 = NewError(1006, "sign error")
	ErrInternalError             = NewError(1007, "internal server error")
	ErrotModifyGroupProperty     = NewError(1008, "cannot modify group property")
	ErrExceedCPS                 = NewError(1309, "exceeds allowed CPS")
	ErrAlreadyExists             = NewError(1530, "already exists")
	ErrNotFound                  = NewError(1531, "not found")
	ErrGroupNotFound             = NewError(1551, "group not found")
	ErrGroupWriteForbidden       = NewError(1553, "group write forbidden")
	ErrOnlySupportSingleFaceMode = NewError(1562, "only support single face mode")
	ErrPoorPictureQuality        = NewError(1563, "poor picture quality")
)

// this map has to be consistent with error variables defined above
var ErrMap = map[int32]*Error{
	1001: ErrMethodNotSupport,
	1002: ErrInvalidParam,
	1003: ErrPermissionDenied,
	1004: ErrInvalidToken,
	1005: ErrExpired,
	1006: ErrSignError,
	1007: ErrInternalError,
	1008: ErrotModifyGroupProperty,
	1309: ErrExceedCPS,
	1530: ErrAlreadyExists,
	1531: ErrNotFound,
	1551: ErrGroupNotFound,
	1553: ErrGroupWriteForbidden,
	1562: ErrOnlySupportSingleFaceMode,
	1563: ErrPoorPictureQuality,
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d,%s", e.Code, e.Msg)
}
