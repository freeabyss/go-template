package cfgCenter

type CfgInter interface {

	QueryChannelRelationId(req *QueryChannelRelationIdReq) (*QueryChannelRelationIdResp, error)

	QueryCommonBanner(req *QueryCommonBannerReq) (*QueryCommonBannerResp, error)

}
