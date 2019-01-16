package util

import "regexp"

const (
	REGEXP_MOBILE_PHONE  =  `^((13[0-9])|(14[5,7,9])|(15[0-9])|(17[0-9])|(18[0-9])|(19[0-9]))\\d{8}$`

	REGEXP_MAIL  =  `^[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)*@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$`

	REGEXP_CHINESE = `^[\u4e00-\u9fa5]+$`
)


func IsMobilePhone(value string) bool {
	mpRegexp := regexp.MustCompile(REGEXP_MOBILE_PHONE)
	return mpRegexp.MatchString(value)
}

func IsMail(value string) bool {
	mRegexp := regexp.MustCompile(REGEXP_MAIL)
	return mRegexp.MatchString(value)
}

func IsChinese(value string) bool {
	cRegexp := regexp.MustCompile(REGEXP_CHINESE)
	return cRegexp.MatchString(value)
}
