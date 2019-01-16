package data

import (
	. "maizuo.com/back-end/go-template/src/server/data/cfgCenter"
)

type Data struct {
	CfgData
}

func NewData() *Data {
	return &Data{}
}
