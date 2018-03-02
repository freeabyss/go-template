package util

import (
	"time"
	"maizuo.com/back-end/go-template/src/server/common"
	"github.com/spf13/viper"
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
)

type MngMessage struct{
	TimeStamp int64  `json:"timeStamp"`
	BusinessType int `json:"businessType"`
	Data interface{} `json:"data"`
}

func NewMessage(businessType int , data interface{}) *MngMessage {
	return &MngMessage{
		TimeStamp:    time.Now().Local().Unix(),
		BusinessType: businessType,
		Data:         data,
	}
}

func PushMngMessage(message *MngMessage)  {
	loghead := Concat("推送消息msg=",ToString(message))
	defer PanicErrorHandler(loghead)
	topic := viper.GetString("nsq.mng_topic")
	bs,_ := json.Marshal(message)
	err := common.NsqProducer.Publish(topic,bs)
	if err != nil {
		loghead = Concat(loghead,"推送消息异常err=",err.Error())
		fmt.Println(loghead)
		Log(loghead)
	}
}

func PushMng(businessType int , data interface{})  {
	PushMngMessage(NewMessage(businessType,data))
}



func PanicErrorHandler(loghead string) {
	if err := recover(); err != nil {
		msg := fmt.Sprint(loghead, "发生panic异常:err=", errors.Wrap(err, 2).ErrorStack())
		fmt.Println(msg)
		Log(msg)
	}
}