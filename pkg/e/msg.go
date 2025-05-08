package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "Request parameter error",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token has timed out",
	ERROR_AUTH_TOKEN:               "Token generation failed",
	ERROR_AUTH:                     "Token error",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
