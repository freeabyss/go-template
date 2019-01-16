package cfgCenter

import (

	"fmt"
	"github.com/spf13/viper"
	"maizuo.com/back-end/go-template/src/server/util"
	"maizuo.com/back-end/go-template/src/server/errcode"
	"encoding/json"
	"strconv"
)

type CfgData struct{}

var DefaultCfgData = &CfgData{}

func (*CfgData)QueryChannelRelationId(req *QueryChannelRelationIdReq) (*QueryChannelRelationIdResp, error) {
	fmt.Println("进入数据层...")
	resp := &QueryChannelRelationIdResp{}
	domain := viper.GetString("domain.business-cfg")
	url := fmt.Sprintf("%s/api/v1/price-channel/query?",domain)
	params := "agentId="+strconv.Itoa(int(req.AgentId))+"&platformId="+strconv.Itoa(int(req.PlatformId))

	body, err := util.HttpGet(url, params, "")
	fmt.Println("third-response:",body)
	if err != nil {
		resp.Status = errcode.SYSTEM_ERROR.Code
		resp.Msg = errcode.SYSTEM_ERROR.Msg
		return resp, err
	}
	statusMsg, err := errcode.CheckRespStatus(body)
	if err != nil{
		resp.Status = statusMsg.Status
		resp.Msg = statusMsg.Msg
		return resp, err
	}
	if err = json.Unmarshal([]byte(body), resp); err != nil {
		resp.Status = errcode.PARAM_PARSE_ERROR.Code
		resp.Msg = errcode.PARAM_PARSE_ERROR.Msg
		return resp, err
	}
	return resp, nil
}



func (*CfgData)QueryCommonBanner(req *QueryCommonBannerReq) (*QueryCommonBannerResp, error) {
	resp := &QueryCommonBannerResp{}
	domain := viper.GetString("domain.business-cfg")
	url := fmt.Sprintf("%s/api/common-banner/list?", domain)
	params := "bannerType="+strconv.Itoa(int(req.bannerType))+"&agentId="+strconv.Itoa(int(req.agentId))

	body, err := util.HttpGet(url, params, "")
	if err != nil {
		resp.Status = errcode.SYSTEM_ERROR.Code
		resp.Msg = errcode.SYSTEM_ERROR.Msg
		return resp, err
	}
	statusMsg, err := errcode.CheckRespStatus(body)
	if err != nil{
		resp.Status = statusMsg.Status
		resp.Msg = statusMsg.Msg
		return resp, err
	}
	if err = json.Unmarshal([]byte(body), resp); err != nil {
		resp.Status = errcode.PARAM_PARSE_ERROR.Code
		resp.Msg = errcode.PARAM_PARSE_ERROR.Msg
		return resp, err
	}
	return resp, nil
}
