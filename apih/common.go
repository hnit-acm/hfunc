package apih

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Error interface {
	error
	Code() int
}

type Err struct {
	code int
	msg  string
}

func NewErr(code int, msg string) *Err {
	return &Err{code: code, msg: msg}
}

func (e Err) Error() string {
	return e.msg
}

func (e Err) Code() int {
	return e.code
}

const (
	OK = iota
	ERR
)

func JsonResponseOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Msg:  "",
		Code: OK,
		Data: data,
	})
}

func JsonResponseErr(ctx *gin.Context, err error) {
	e, ok := err.(Err)
	code := ERR
	if ok {
		code = e.Code()
	}
	ctx.JSON(http.StatusOK, Response{
		Msg:  err.Error(),
		Code: code,
		Data: nil,
	})
}
