package controller

import (
	"github.com/kataras/iris"
	"maizuo.com/back-end/go-template/src/server/errcode"
	"fmt"
	API "maizuo.com/back-end/go-template/src/server/service/cfg"
)


func doQueryChannelRelationId(ctx *iris.Context, api API.CfgAPI)  {
	loghead := "获取渠道关系ID:"
	agentId, err_a := ctx.URLParamInt("agentId")
	platformId, err_p := ctx.URLParamInt("platformId")
	fmt.Println(fmt.Sprintf(loghead,agentId,platformId))
	if err_a != nil || err_p != nil {
		ctx.JSON(iris.StatusOK, errcode.PARAM_ERROR.Result())
		return
	}
	data, ecode := api.QueryChannelRelationId(uint32(agentId), uint32(platformId))
	fmt.Println("ecode:",ecode)
	if ecode != nil {
		ctx.JSON(iris.StatusOK, ecode.Result())
		return
	}
	ctx.JSON(iris.StatusOK, errcode.SUCCESS.ResultWithData(data))
}