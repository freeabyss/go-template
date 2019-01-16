package initialize

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/kataras/iris"
	"maizuo.com/back-end/go-template/src/server/errcode"
	"maizuo.com/back-end/go-template/src/server/util"
)

func SetErrorDeal() {
	iris.Use(iris.HandlerFunc(func(ctx *iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("发生panic异常: %v\n", errors.Wrap(err, 2).ErrorStack())
				util.LogInfo(ctx, errcode.SYSTEM_ERROR.ResultWithMsg(msg))
				ctx.JSON(iris.StatusOK, errcode.SYSTEM_ERROR.Result())
			}
		}()
		ctx.Next()
	}))

}
