package errno

import "net/http"

var (
	OK                  = &Errno{HTTPStatus: http.StatusOK, Code: 0, Msg: "ok"}
	InternalServerError = &Errno{HTTPStatus: http.StatusInternalServerError, Code: 1, Msg: "internal server error"}
	ErrParam            = &Errno{HTTPStatus: http.StatusBadRequest, Code: 2, Msg: "request param error"}
	ErrTokenInvalid     = &Errno{HTTPStatus: http.StatusUnauthorized, Code: 3, Msg: "the token is invalid"}
	ErrAccount          = &Errno{HTTPStatus: http.StatusUnauthorized, Code: 4, Msg: "account or password error"}
	ErrPassword         = &Errno{HTTPStatus: http.StatusBadRequest, Code: 5, Msg: "new password must contain at least 8 characters"}
	ErrDataExist        = &Errno{HTTPStatus: http.StatusConflict, Code: 7, Msg: "data already exists"}
	ErrDataNotExist     = &Errno{HTTPStatus: http.StatusNotFound, Code: 8, Msg: "data does not exist"}
)
