package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"maizuo.com/back-end/go-template/src/server/entity"
	"maizuo.com/back-end/go-template/src/server/util"
)

/**
基础controller
*/
type BaseController struct {
}

/**
日志记录
*/
func (*BaseController) SetLog(ctx *iris.Context) {
	ctx.Next()
	//接口返回自动写入日志
	var result entity.Result
	json.Unmarshal(ctx.Response.Body(), &result)
	util.LogInfo(ctx, &result)
}

//func (*BaseController) HeartCheck(ctx *iris.Context) {
//	ctx.JSON(iris.StatusOK, heart_check.DefaultHeartCheckService.HeartCheck())
//}
