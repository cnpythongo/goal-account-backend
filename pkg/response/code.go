package response

import (
	"github.com/cnpythongo/goal/pkg/common/config"
)

const (
	SuccessCode  = 0
	FailCode     = 1
	UnknownError = 99999

	JWTTokenEmptyError   = 10000
	JWTTokenParseError   = 10001
	JWTTokenExpiredError = 10002
	JWTTokenGenError     = 10003

	DbError = 11000

	PayloadError          = 12000
	PayloadParseJsonError = 12001
	PayloadUnmarshalError = 12002
	PayloadCopyError      = 12003

	AccountLoginError          = 13000
	AccountEmailExistsError    = 13001
	AccountCreateError         = 13002
	AccountUserIdError         = 13003
	AccountUserNotExistError   = 13004
	AccountQueryUserError      = 13005
	AccountQueryUserParamError = 13006
	AccountQueryUserListError  = 13007
	AccountPasswordError       = 13008
	AccountUserExistError      = 13009
	AccountUserUpdateError     = 13010
	AccountUserFreezeError     = 13011
	AccountUserDeleteError     = 13012
)

var MsgMapping = map[string]map[int]string{
	"en":    MessageEn,
	"zh_cn": MessageZHCN,
}

func GetCodeMsg(code int) string {
	lang := config.GetConfig().App.Language
	mapping, ok := MsgMapping[lang]
	if !ok {
		return ""
	}
	msg, ok := mapping[code]
	if !ok {
		return ""
	}
	return msg
}
