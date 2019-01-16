package util

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris"
	"gopkg.in/square/go-jose.v1/json"
	"errors"
)

func GetHeaderData (ctx *iris.Context) (error, string ){
	client := ctx.RequestHeader("X-client-info")
	if client == `` {
		client = `{}`
	}
	fmt.Println("x-client-into", client)
	clientInfo := make(map[string]interface{})
	//clientInfo := ClientInfo{}
	rClient,_ := simplejson.NewJson([]byte(client))

	token := ctx.RequestHeader("x-token")
	if token == `` {
		fmt.Println("没有x-token")
		token = `{}`
	}
	fmt.Println("x-token", token)
	rToken, _ := simplejson.NewJson([]byte(token))

	clientInfo["appStoreId"] = rClient.Get("a")
	clientInfo["carrier"] = rClient.Get("ca")
	clientInfo["channelId"] = rClient.Get("ch")
	clientInfo["cityId"] = rClient.Get("ci")
	clientInfo["equipmentId"] = rClient.Get("e")
	clientInfo["clientIp"] = rClient.Get("i")
	clientInfo["latitude"] = rClient.Get("la")
	clientInfo["longitude"] = rClient.Get("lo")
	clientInfo["locationType"] = rClient.Get("l")
	clientInfo["version"] = rClient.Get("v")
	clientInfo["userId"] = rToken.Get("userId")

	cInfo, err := json.Marshal(clientInfo)

	if err != nil {
		fmt.Println("转换失败2————————》", err)
		return errors.New("转换失败"), ""
	}
	//fmt.Println("x-client-into2", string(cInfo))
	clientInfoBase64 := Base64(string(cInfo))
	return nil, clientInfoBase64
}
