package cfg

import (
	"maizuo.com/back-end/go-template/src/server/data/cfgCenter"
	"maizuo.com/back-end/go-template/src/server/errcode"
	"maizuo.com/back-end/go-template/src/server/data"
	"fmt"
)

type CfgService struct {}

var (
	api cfgCenter.CfgInter = data.NewData()
)

func (*CfgService)QueryChannelRelationId(agentId, platformId uint32) (*QueryChannelRelationIdRespData, *errcode.ErrCode) {
	fmt.Println("apisllllll", api)
	thirdReq := &cfgCenter.QueryChannelRelationIdReq{AgentId:agentId, PlatformId:platformId}
	thirdResp,err := api.QueryChannelRelationId(thirdReq)
	if err != nil{
		return nil,&errcode.ErrCode{Code:thirdResp.Status, Msg:thirdResp.Msg}
	}
	resp := &QueryChannelRelationIdRespData{}
	resp.ChannelId = thirdResp.Data.ChannelId
	return resp, nil
}
