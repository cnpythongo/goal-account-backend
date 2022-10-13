package response

var MessageZHCN = map[int]string{
	SuccessCode:  "ok",
	FailCode:     "失败",
	UnknownError: "未知错误",

	JWTTokenEmptyError:   "未提供token",
	JWTTokenParseError:   "会话token解析失败",
	JWTTokenExpiredError: "会话已过期，请重新登录",
	JWTTokenGenError:     "生成会话Token失败",

	DbError: "数据库操作失败",

	PayloadError:          "请求的数据内容不正确",
	PayloadParseJsonError: "JSON解析失败",
	PayloadUnmarshalError: "对象转换失败",
	PayloadCopyError:      "对象属性复制失败",

	AccountLoginError:          "登录失败",
	AccountEmailExistsError:    "邮箱地址已存在，请换一个",
	AccountCreateError:         "创建用户失败",
	AccountUserIdError:         "用户id不正确",
	AccountUserNotExistError:   "用户不存在",
	AccountQueryUserError:      "查询用户失败",
	AccountQueryUserParamError: "查询用户参数不正确",
	AccountQueryUserListError:  "查询用户列表数据失败",
	AccountPasswordError:       "密码不正确",
	AccountUserExistError:      "用户名已存在，请换一个",
	AccountUserUpdateError:     "更新用户数据失败",
	AccountUserFreezeError:     "该用户已被冻结",
	AccountUserDeleteError:     "删除用户数据失败",
}
