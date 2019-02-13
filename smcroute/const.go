package smcroute

const (
	NULL_CHARACTER        byte   = byte(0x00)
	NULL_CHARACTER_STRING string = "\x00"

	RESPONSE_BUFFER_SIZE int = 255

	DEFAULT_SOCKET_PATH string = "/run/smcroute.sock"
	DEFAULT_NETWORK     string = "unix"
)
