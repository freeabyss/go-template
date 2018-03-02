package controller

import (
	"maizuo.com/back-end/go-template/src/server/service"
	"github.com/kataras/iris"
	"fmt"
)
type HandleController struct {}

var (
	api = service.NewService()
)


func (h *HandleController) QueryChannelRelationId(ctx *iris.Context) {
	fmt.Println("api:",api)
	doQueryChannelRelationId(ctx, api)
}