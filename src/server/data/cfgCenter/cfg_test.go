package cfgCenter

import (
	"testing"
	"fmt"
	"maizuo.com/back-end/go-template/src/server/util"
)

func TestQueryChannelRelationId(t *testing.T)  {
	//initialize.SetupLogger()
	//var cfgData = &CfgData{}
	util.InitConfigForTest(util.GetRelToConfigPath("local"))
	util.SetUpLoggerForTest()
	resp, err := DefaultCfgData.QueryChannelRelationId(&QueryChannelRelationIdReq{
		AgentId:uint32(1000),
		PlatformId:uint32(2),
	})
	if err != nil {
		fmt.Println("cfgData.QueryChannelRelationId err=",err)
	}
	fmt.Println("cfgData.QueryChannelRelationId resp=",resp)
}
