package app

import (
	"github.com/gin-gonic/gin"
	app_err "github.com/nnnn24/go-gin-boilerplate/pkg/err"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  app_err.GetMsg(errCode),
		Data: data,
	})

	return
}
