package errno

var (
	OK                  = &Errno{Code: 0, Msg: "ok"}
	InternalServerError = &Errno{Code: 1, Msg: "internal server error"}
	ErrParam            = &Errno{Code: 2, Msg: "request param error"}
	ErrTokenInvalid     = &Errno{Code: 3, Msg: "the token was invalid"}
	ErrAccount          = &Errno{Code: 4, Msg: "account or password error"}
	ErrAccountPwdHasHan = &Errno{Code: 5, Msg: "pwd has han"}
	ErrAccountNewPwdLen = &Errno{Code: 6, Msg: "the len of new pass should >= 8"}
	ErrDataExist        = &Errno{Code: 7, Msg: "data existed"}
	ErrDataNotExist     = &Errno{Code: 8, Msg: "data not exist"}
)
