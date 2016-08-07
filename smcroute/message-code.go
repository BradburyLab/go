package smcroute

import "fmt"

type MessageCode uint16

const ( // base
	MESSAGE_UNKNOWN MessageCode = 0

	INFO  MessageCode = 0x8000
	ERROR MessageCode = 0x4000
)

const ( // reserved for future use
	_ = iota // skip; now iota == 1

	INFO_OK_JOIN MessageCode = INFO | iota
	INFO_OK_LEAVE
)

const (
	_ = iota // skip; now iota = 1

	ERROR_SOCKET_CONNECT MessageCode = ERROR | iota
	ERROR_SOCKET_WRITE
	ERROR_SOCKET_READ
	ERROR_CMD_ENCODE
	ERROR_EXEC
	ERROR_DROP_MEMBERSHIP_FAILED_99
)

var messageCodeText = map[MessageCode]string{
	ERROR_SOCKET_CONNECT: `error connecting to smcroute daemon: {"socket-path": "%s", "error":"%s"};`,
	ERROR_SOCKET_WRITE:   `error writing to smcroute daemon: {"cmd-bash": "%s", "socket-path": "%s", "error": "%s"};`,
	ERROR_SOCKET_READ:    `error reading from smcroute daemon: {"cmd-bash": "%s", "socket-path": "%s", "error": "%s"};`,
	ERROR_CMD_ENCODE:     `error encoding cmd into byte array: {"cmd-bash": "%s", "error": "%s"};`,
	ERROR_EXEC:           `error executing cmd, see error string from smcroute daemon: {"cmd-bash": "%s", "latency": "%s", "error": "%s"};`,

	// leave => no routes was assigned, nothing to leave
	ERROR_DROP_MEMBERSHIP_FAILED_99: `DROP MEMBERSHIP failed. Error 99: Cannot assign requested address`,
}

var messageCodeString = map[MessageCode]string{
	ERROR_SOCKET_CONNECT: "error-socket-connect",
	ERROR_SOCKET_WRITE:   "error-socket-write",
	ERROR_SOCKET_READ:    "error-socket-read",
	ERROR_CMD_ENCODE:     "error-cmd-encode",
	ERROR_EXEC:           "error-exec",
}

func (it MessageCode) String() string { return MessageCodeString(it) }

func MessageCodeText(it MessageCode) string {
	if v, ok := messageCodeText[it]; ok {
		return v
	}
	return fmt.Sprintf("missing message-text for 0x%X message", uint16(it))
}

func MessageCodeString(it MessageCode) string {
	if v, ok := messageCodeString[it]; ok {
		return v
	}
	return fmt.Sprintf("missing message-string for 0x%X message", uint16(it))
}
