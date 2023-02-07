package e

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALID_PARAMS:          "Invalid params",
	ERROR_EXIST_TAG:         "Tag already exists",
	ERROR_NOT_EXIST_TAG:     "Tag not exist",
	ERROR_NOT_EXIST_ARTICLE: "Article not exist",

	ERROR_EXIST_PAGE:     "Page already exists",
	ERROR_NOT_EXIST_PAGE: "Page not exist",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Auth failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token timed out",
	ERROR_AUTH_TOKEN:               "Token Generation Failed",
	ERROR_AUTH:                     "Invalid Token",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
