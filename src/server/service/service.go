package service

import (
	. "maizuo.com/back-end/go-template/src/server/service/cfg"
)
type Service struct {
	CfgService
}

func NewService() *Service {
	return &Service{}
}
