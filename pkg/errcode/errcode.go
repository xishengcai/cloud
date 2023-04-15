package errcode

import (
	"fmt"
	"net/http"

	"github.com/aagu/go-i18n/pkg/translation"
)

type Error struct {
	code     int64
	msg      *translation.Message
	template interface{}
}

var codes = map[int64]*translation.Message{}

func NewError(code int64, msg *translation.Message) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d already exist, please use another one", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) clone() *Error {
	return &Error{
		code:     e.code,
		msg:      e.msg,
		template: e.template,
	}
}

func (e *Error) Error() string {
	return e.Msg()
}

func (e *Error) Code() int64 {
	return e.code
}

func (e *Error) Msg() string {
	if e.template != nil {
		return e.msg.Format(e.template)
	}
	return e.msg.String()
}

func (e *Error) Msgf(v interface{}) string {
	return e.msg.Format(v)
}

func (e *Error) TranslatableMsg() *translation.Message {
	return e.msg
}

func (e *Error) TranslatableTemplate() interface{} {
	return e.template
}

func (e *Error) Format(v interface{}) *Error {
	clone := e.clone()
	clone.template = v
	return clone
}

func (e *Error) HttpCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case BadRequest.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case TokenNotFound.Code():
		return http.StatusUnauthorized
	case TokenValidError.Code():
		return http.StatusUnauthorized
	case NoPermission.Code():
		return http.StatusUnauthorized
	}

	return http.StatusOK
}

func HttpError(code int) error {
	switch code {
	case http.StatusNotFound:
		return NotFound
	case http.StatusBadRequest:
		return InvalidParams
	case http.StatusConflict:
		return Conflict
	case http.StatusUnauthorized:
		return NoPermission.Format("")
	case http.StatusServiceUnavailable:
		return Unavailable
	case http.StatusForbidden:
		return Unsupported
	case http.StatusInternalServerError:
		return ServerError
	}
	return nil
}

func AsErrCode(e error) (*Error, bool) {
	ecode, ok := e.(*Error)
	return ecode, ok
}
