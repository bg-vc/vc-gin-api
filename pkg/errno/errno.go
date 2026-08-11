package errno

import "errors"

type Errno struct {
	HTTPStatus int
	Code       int
	Msg        string
}

func (err *Errno) Error() string {
	return err.Msg
}

func Decode(err error) (int, int, string) {
	if err == nil {
		return OK.HTTPStatus, OK.Code, OK.Msg
	}
	var known *Errno
	if errors.As(err, &known) {
		return known.HTTPStatus, known.Code, known.Msg
	}
	return InternalServerError.HTTPStatus, InternalServerError.Code, InternalServerError.Msg
}
