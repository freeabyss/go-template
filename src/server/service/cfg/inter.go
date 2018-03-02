package cfg

import "maizuo.com/back-end/go-template/src/server/errcode"

type CfgAPI interface {
	QueryChannelRelationId(agentId, platformId uint32) (*QueryChannelRelationIdRespData, *errcode.ErrCode)
}