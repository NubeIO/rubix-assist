package response

var MsgFlags = map[int]string{
	SUCCESS:        "Ok",
	INVALID_PARAMS: "Invalid params error",
	ERROR:          "Fail",
	FILENOTEXIST:   "File doesn't exist",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
