package hapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg    string      `json:"msg"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Status int         `json:"-"`
}

type Error interface {
	error
	Code() int
}

type withCodeErr struct {
	code int
	msg  string
}

func NewCodeErr(code int, msg string) withCodeErr {
	return withCodeErr{code: code, msg: msg}
}

func (e withCodeErr) Error() string {
	return e.msg
}

const (
	OK = iota
	ERR
)

type Provide func(Response) Response

func WithStatus(status int) Provide {
	return func(res Response) Response {
		res.Status = status
		return res
	}
}

func WithData(data interface{}) Provide {
	return func(res Response) Response {
		res.Data = data
		return res
	}
}

func WithCode(code int) Provide {
	return func(res Response) Response {
		res.Code = code
		return res
	}
}

func WithMsg(msg string) Provide {
	return func(res Response) Response {
		res.Msg = msg
		return res
	}
}

func WithErr(err error) Provide {
	return func(res Response) Response {
		res.Msg = err.Error()
		e, ok := err.(withCodeErr)
		if ok {
			res.Code = e.code
			res.Msg = e.msg
		}
		return res
	}
}

func JsonResponse(ctx *gin.Context, res Response, p ...Provide) {
	for _, provide := range p {
		res = provide(res)
	}
	ctx.JSON(res.Status, res)
}

func JsonResponseOk(ctx *gin.Context, data interface{}, p ...Provide) {
	res := Response{
		Status: http.StatusOK,
		Code:   OK,
		Data:   data,
	}
	JsonResponse(ctx, res, p...)
}

func JsonResponseErr(ctx *gin.Context, err error, p ...Provide) {
	res := Response{
		Status: http.StatusOK,
		Code:   ERR,
		Msg:    err.Error(),
	}
	e, ok := err.(withCodeErr)
	if ok {
		res.Code = e.code
		res.Msg = e.msg
	}
	JsonResponse(ctx, res, p...)
}
