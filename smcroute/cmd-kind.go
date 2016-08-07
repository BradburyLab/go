package smcroute

type CMDKind uint16

const (
	CMD_UNKNOWN CMDKind = 0
	CMD_J       CMDKind = 'j' // Join a multicast group
	CMD_L       CMDKind = 'l' // Leave a multicast group
	CMD_A       CMDKind = 'a' // Add a multicast route
	CMD_R       CMDKind = 'r' // Remove a multicast route

	CMD_JOIN   = CMD_J
	CMD_LEAVE  = CMD_L
	CMD_ADD    = CMD_A
	CMD_REMOVE = CMD_R
)

var cmdKindString = map[CMDKind]string{
	CMD_UNKNOWN: "unknown-smcroute-command",
	CMD_JOIN:    "join",
	CMD_LEAVE:   "leave",
	CMD_ADD:     "add",
	CMD_REMOVE:  "remove",
}

func (it CMDKind) String() string { return cmdKindString[it] }
