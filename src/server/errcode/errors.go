package errcode

import (
	"maizuo.com/back-end/go-template/src/server/entity"
	"github.com/bitly/go-simplejson"
	"errors"
)

//返回结果错误码定义
var (
	SUCCESS = &ErrCode{0, "success"}
	SYSTEM_ERROR = &ErrCode{-1, "系统异常"}
	PARAM_ERROR = &ErrCode{1005, "请求参数非法"}
	PARAM_PARSE_ERROR = &ErrCode{1007, "json或xml解析错误"}
)

type ErrCode struct {
	Code int
	Msg  string
}

func (e *ErrCode)Result() *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:e.Msg,
		Data:"",
	}
}

func (e *ErrCode)ResultWithData(data interface{}) *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:e.Msg,
		Data:data,
	}
}

func (e *ErrCode)ResultWithMsg(msg string) *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:msg,
		Data:"",
	}
}

func  (e *ErrCode)ReplaceMsg(msg string) *ErrCode {
	return &ErrCode{
		Code:e.Code,
		Msg:msg,
	}
}

// check response status
func CheckRespStatus(body string) (statusMsg *entity.StatusMsg, err error) {
	statusMsg = &entity.StatusMsg{}
	resultJson, err := simplejson.NewJson([]byte(body))
	if err != nil { //json数据转换失败
		statusMsg.Status = PARAM_PARSE_ERROR.Code
		statusMsg.Msg = PARAM_PARSE_ERROR.Msg
		return
	}
	if resultJson.Get("status").MustInt() != 0 {
		statusMsg.Status = resultJson.Get("status").MustInt()
		statusMsg.Msg = resultJson.Get("msg").MustString()
		err = errors.New(resultJson.Get("msg").MustString())
		return
	}
	return
}

