package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Sirupsen/logrus"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
	"maizuo.com/back-end/go-template/src/server/common"
	"maizuo.com/back-end/go-template/src/server/entity"
)

var LogInfo = func(ctx *iris.Context, result *entity.Result) {
	var processTime int64
	startAt := ctx.Get("startAt")
	if startAt != nil {
		startAt := startAt.(int64)
		endAt := time.Now().UnixNano() / 1000000
		processTime = endAt - startAt
	} else {
		processTime = -1
	}

	body := string(ctx.PostBody()[:])

	if result == nil {
		result = &entity.Result{0, nil, "success"}
	}
	var data interface{}
	alarmID := "0"
	handle := strings.Split(ctx.GetHandlerName(), "/")
	var _api []string
	var _f []string
	var _c string = ""
	var _system string = ""
	var _controller string = ""
	var _func string = ""
	var _interface string = ""
	if len(handle) >= 2 {
		_api = strings.Split(handle[len(handle)-1], ".")
		if len(_api) >= 2 {
			_f = strings.Split(_api[len(_api)-1], "-")
			_c = _api[len(_api)-2]
		}
		if len(_c) >= 3 {
			_controller = _c[2 : len(_c)-1]
		}
		_system = strings.Replace(handle[1], "-", "_", -1)
		if len(_f) >= 1 {
			_func = _f[0]
		}
	}
	_interface = _system + "_" + _controller + "_" + _func
	if result.Status != 0 {
		_interface = "error_" + _interface
		alarmID = "1"
		data = result.Data
	}
	Logger := logrus.WithFields(logrus.Fields{
		"@source":    ctx.LocalAddr().String(),
		"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
		"@fields": map[string]interface{}{
			"fromtype":    viper.GetString("name"),
			"host":        ctx.HostString(),
			"interface":   _interface,
			"method":      ctx.MethodString(),
			"ip":          ctx.RemoteAddr(),
			"query":       fmt.Sprint(ctx.URLParams()),
			"param":       fmt.Sprint(ctx.ParamsSentence()),
			"body":        body,
			"alarmID":     alarmID,
			"path":        ctx.PathString(),
			"processTime": strconv.FormatInt(processTime, 10),
			"msg":         result.Msg,
			"data":        data,
			"status":      strconv.Itoa(result.Status),
			"system":      viper.GetString("name"),
			"totype":      viper.GetString("name"),
			"errorType":   strconv.Itoa(result.Status),
			"other":       result.Msg,
		},
	})
	Logger.Warningln(result.Status)
}

func AlarmLog(interfaceName, param, other string) {
	common.Logger.WithFields(
		logrus.Fields{
			"@source":    GetLocalIP(),
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"fromtype":    "72",
				"interface":   interfaceName,
				"param":       param,
				"result":      "1",
				"alarmID":     "1",
				"processTime": "0",
				"other":       other,
				"system":      viper.GetString("name"),
				"totype":      "72",
				"errorType":   "1",
			},
		}).Warningln()
}

func Log(params ... interface{}) {
	var msg =fmt.Sprintln(params)
	common.Logger.WithFields(
		logrus.Fields{
			"@source":    GetLocalIP(),
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"system": viper.GetString("name"),
				"msg":    msg,
			},
		}).Warningln()
}
