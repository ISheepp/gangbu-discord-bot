package e

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "fail",
	DiscordUserNotFound:       "discord user not found",
	ErrorDiscordId:            "discord id error, may not snowflake",
	ErrorEthereumNodeResponse: "eth node response error",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
