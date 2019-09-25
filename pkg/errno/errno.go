package errno

import "fmt"

type Errno struct {
	Code int
	Msg  string
}

func (err Errno) Error() string {
	return err.Msg
}

type Err struct {
	Code int
	Msg  string
	Err  error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, msg: %s, error: %s", err.Code, err.Msg, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Msg
	case *Errno:
		return typed.Code, typed.Msg
	default:
	}
	return InternalServerError.Code, err.Error()
}
