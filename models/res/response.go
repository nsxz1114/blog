package res

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/utils"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type ListResponse[T any] struct {
	Count int `json:"count"`
	List  T   `json:"list"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func Ok(c *gin.Context) {
	Result(Success, map[string]any{}, "success", c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "success", c)
}
func OkWithList(list any, count int, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}
func OkWithCode(code ResCode, c *gin.Context) {
	Result(Success, map[string]any{}, code.Msg(), c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
func FailWithCode(code ResCode, c *gin.Context) {
	Result(Error, map[string]any{}, code.Msg(), c)
}
