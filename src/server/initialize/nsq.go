package initialize

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"maizuo.com/back-end/go-template/src/server/common"
)

func SetupNsqProducer() {
	tcp := viper.GetString("nsq.tcp")
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(tcp, config)
	if err != nil {
		msg := fmt.Sprint("failed to connect nsq,tcp=", tcp, "err=", err.Error())
		panic(msg)
	}
	//  try to ping
	err = producer.Ping()
	if nil != err {
		producer.Stop()
		producer = nil
		msg := fmt.Sprint("failed to connect nsq,tcp=", tcp, "err=", err.Error())
		panic(msg)
	}
	common.NsqProducer = producer
}
