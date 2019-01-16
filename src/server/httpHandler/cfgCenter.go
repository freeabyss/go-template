package httpHandler

import (

	"maizuo.com/back-end/app-cfg/src/server/entity/topic"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"maizuo.com/back-end/app-cfg/src/server/util"
	"maizuo.com/back-end/app-cfg/src/server/errcode"
	"encoding/json"
	"maizuo.com/back-end/app-cfg/src/server/entity"
)

func QueryPriceChannelId(req *topic.PriceChannelReq) (resp *topic.PriceChannelResp, err error) {
	resp = &topic.PriceChannelResp{}
	domain := viper.GetString("domain.business-cfg")
	url := fmt.Sprintf("%s/api/v1/price-channel/query?",domain)
	params := "agentId="+strconv.Itoa(int(req.AgentId))+"&platformId="+strconv.Itoa(int(req.PlatformId))

	body, err := util.HttpGet(url, params, "")
	if err != nil {
		resp.Status = errcode.SYSTEM_ERROR.Code
		resp.Msg = errcode.SYSTEM_ERROR.Msg
		return
	}
	statusMsg, err := errcode.CheckRespStatus(body)
	if err != nil{
		resp.Status = statusMsg.Status
		resp.Msg = statusMsg.Msg
		return
	}
	if err = json.Unmarshal([]byte(body), resp); err != nil {
		resp.Status = errcode.PARAM_PARSE_ERROR.Code
		resp.Msg = errcode.PARAM_PARSE_ERROR.Msg
		return
	}
	return
}



func QueryCommonBannerV2(bannerType, agentId uint32) (resp *entity.BannerResult, err error) {
	resp = &entity.BannerResult{}
	domain := viper.GetString("domain.business-cfg")
	url := fmt.Sprintf("%s/api/common-banner/list?", domain)
	params := "bannerType="+strconv.Itoa(int(bannerType))+"&agentId="+strconv.Itoa(int(agentId))

	body, err := util.HttpGet(url, params, "")
	if err != nil {
		resp.Status = errcode.SYSTEM_ERROR.Code
		resp.Msg = errcode.SYSTEM_ERROR.Msg
		return
	}
	statusMsg, err := errcode.CheckRespStatus(body)
	if err != nil{
		resp.Status = statusMsg.Status
		resp.Msg = statusMsg.Msg
		return
	}
	if err = json.Unmarshal([]byte(body), resp); err != nil {
		resp.Status = errcode.PARAM_PARSE_ERROR.Code
		resp.Msg = errcode.PARAM_PARSE_ERROR.Msg
		return
	}
	return
}
