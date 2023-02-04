package e

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALID_PARAMS:          "Invalid params",
	ERROR_EXIST_TAG:         "Tag already exists",
	ERROR_NOT_EXIST_TAG:     "Tag not exist",
	ERROR_NOT_EXIST_ARTICLE: "Article not exist",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}