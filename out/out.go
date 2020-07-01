package out

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct{
	Code	int				`json:"code"`
	Msg		string			`json:"msg"`
	Data	interface{}		`json:"data"`
}

//成功
func ReturnSuccess(c *gin.Context,obj interface{})  {
	result := Result{
		Code: 0,
		Msg: "",
		Data: obj,
	}
	c.JSON(http.StatusOK, result)
}

//失败
func ReturnFail(c *gin.Context,code int,msg string)  {
	result := Result{
		Code: code,
		Msg: msg,
	}
	c.JSON(http.StatusInternalServerError, result)
}

//授权信息失效
func ReturnUnauthorized(c *gin.Context,code int,msg string)  {
	result := Result{
		Code: code,
		Msg: msg,
	}
	c.JSON(http.StatusUnauthorized, result)
}

//服务器拒绝访问
func ReturnForbidden(c *gin.Context,code int,msg string)  {
	result := Result{
		Code: code,
		Msg: msg,
	}
	c.JSON(http.StatusForbidden, result)
}


//返回Json
func ReturnJson(c *gin.Context,code int,result Result)  {
	c.JSON(result.Code, result)
}
