package util

import (
	"maizuo.com/back-end/app-cfg/src/server/common"
	"github.com/spf13/viper"
	"errors"
	"time"
)

func SetNX(businessName, uniqueKey string, second int64) bool {
	key := businessName + ":" + uniqueKey
	rs := common.Redis.SetNX(key, "1", time.Duration(int64(time.Second) * second))
	if rs.Val() == true {
		return true
	}
	return false
}

func RateLimit(businessName, uniqueKey string, count, second int64, needAdd bool) (bool, error) {
	key := viper.GetString("rateLimit.key") + "_" + businessName + "_" + uniqueKey
	if key == "" {
		return false, errors.New("取不到reids中ratelimit")
	}
	//获取当前队列长度
	len := common.Redis.LLen(key).Val()
	//如果请求次数已经超过当前最大的请求次数
	if len >= count {
		firstTime, err := common.Redis.RPop(key).Int64()
		if err != nil {
			return false, errors.New("取不到reids中ratelimitTime")
		}
		//判断当前时间是否已经超过限制时间,恢复次数
		if time.Now().Unix() > firstTime + second {
			//队列push 将当前请求时间放入队列中
			if needAdd {
				common.Redis.LPush(key, time.Now().Unix())
				//设置过期时间
				common.Redis.Expire(key, time.Duration(second * int64(time.Second)))
			}
			return true, nil
		} else {
			common.Redis.RPush(key, firstTime)
			return false, nil
		}

	}

	//队列push 将当前请求时间放入队列中
	if needAdd {
		common.Redis.LPush(key, time.Now().Unix())
		common.Redis.Expire(key, time.Duration(second * int64(time.Second)))
	}

	return true, nil
}