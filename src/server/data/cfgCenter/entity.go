package cfgCenter


type QueryChannelRelationIdReq struct {
	AgentId 	uint32         `json:"agentId"`
	PlatformId    	uint32      `json:"platformId"`
}

type QueryChannelRelationIdResp struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   struct{
		       ChannelId	uint32	  `json:"channelId"`
	       } `json:"data"`
}


type QueryCommonBannerReq struct {
	bannerType 	uint32      `json:"bannerType"`
	agentId    	uint32      `json:"agentId"`
}

type QueryCommonBannerResp struct {
	BannerList 		[]*Banner 		`json:"data"`
	Msg    			string 			`json:"msg"`
	Status 			int  			`json:"status"`
}

type Banner struct {
	BannerId   	uint64 `json:"bannerId"`
	ImgUrl     	string `json:"imgUrl"`
	ActionType 	uint32 `json:"actionType"`
	ActionData 	string `json:"actionData"`
	MasterTitle     string `json:"masterTitle"`
	SlaveTitle 	string `json:"slaveTitle"`
}