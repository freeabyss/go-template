package util

import (
	"errors"
	"github.com/parnurzeal/gorequest"
)


func HttpGet(url string, params interface{}, header string) (string, error) {

	request := gorequest.New()
	resp, body, errs := request.Get(url).Set("malldata",header).Query(params).End()
	for _, err := range errs {
		if err != nil {
			Log("get请求链接地址:", url, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	Log("请求链接地址:", resp.Request.URL, params, "返回的结果是: ", body)
	return body, nil
}

func HttpPostForm(url string, params interface{}) (string, error) {

	request := gorequest.New()
	resp, body, errs := request.Post(url).Type("multipart").Send(params).End()
	for _, err := range errs {
		if err != nil {
			Log("postForm请求链接地址:", url, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	Log("请求链接地址:", resp.Request.URL, params, "返回的结果是: ", body)
	return body, nil
}

func HttpPostJson(url string, params interface{}, header string) (string, error) {

	request := gorequest.New()
	resp, body, errs := request.Post(url).Set("malldata",header).Send(params).End()
	for _, err := range errs {
		if err != nil {
			Log("postJson请求链接地址:",url, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	Log("请求链接地址:", resp.Request.URL, params, "返回的结果是: ", body)
	return body, nil
}

