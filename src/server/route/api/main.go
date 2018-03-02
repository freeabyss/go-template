package api

import (
	"github.com/kataras/iris"
	"maizuo.com/back-end/go-template/src/server/controller"
)

func Api() {

	var (
		base = &controller.BaseController{}
		handle = &controller.HandleController{}
	)
	//日志记录
	api := iris.Party("/api", base.SetLog)
	//api.Get("/check/heartbeat/","")
	api.Get("/channel-relation-id/query/", handle.QueryChannelRelationId)
 }
