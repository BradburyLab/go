package smcroute

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"unsafe"
)

type header struct {
	length    int64
	cmd       CMDKind
	argsCount uint16
}

type CMD struct {
	header

	args     []byte
	argsOrig []string
}

func (it *CMD) StringBash() string {
	return fmt.Sprintf(
		"sudo smcroute -%c %s",
		it.cmd,
		strings.Join(it.argsOrig, " "),
	)
}

// -j eth0.33 239.255.11.101
//
//  +----+-----+---+-------------------------------+
//  | 40 | 'j' | 2 | "eth0.33\0239.255.11.101\0\0" |
//  +----+-----+---+-------------------------------+
//  ^              ^
//  |              |
//  |              |
//  +-----cmd------+
//
//  sizeof(struct cmd) = 16
//  strlen(args) = 21
//  sizeof(3 NULL_CHARACTERS) = 3
//    2 => after each string arument
//    1 => at the end of arfs bye array
//
//  length = 16 + 21 + 3 = 40 bytes
//
//  strace: write(3, "(\0\0\0\0\0\0\0j\0\2\0\0\0\0\0eth0.33\000239.255.11.101\0\0", 40) = 40
func (it *CMD) Encode() (*bytes.Buffer, error) {
	b := make([]byte, it.length)
	buf := bytes.NewBuffer(b)
	buf.Reset()

	if e := binary.Write(buf, binary.LittleEndian, it.header); e != nil {
		return nil, e
	}

	// add 0 padding
	// e.g. 1. sizeof(struct{size_t, uint16, uint16}) in C is 16 bytes
	//      2. sizeof(size_t) + sizeof(uint16) + sizeof(uint16) = 8 + 2 + 2 = 12 bytes
	//      3. 16 - 12 = 4 bytes of padding should be added
	structPaddedSize := int(unsafe.Sizeof(it.header))
	for structPaddedSize-buf.Len() > 0 {
		if e := buf.WriteByte(NULL_CHARACTER); e != nil {
			return nil, e
		}
	}

	if _, e := buf.Write(it.args); e != nil {
		return nil, e
	}

	return buf, nil
}

func NewCMD(cmd CMDKind, args ...string) *CMD {
	it := new(CMD)

	it.header.cmd = cmd
	it.header.argsCount = uint16(len(args))

	for _, arg := range args {
		it.args = append(it.args, []byte(arg)...)
		it.args = append(it.args, NULL_CHARACTER)
	}
	it.args = append(it.args, NULL_CHARACTER)
	it.argsOrig = args

	it.header.length = int64(unsafe.Sizeof(it.header)) + int64(len(it.args))

	return it
}
