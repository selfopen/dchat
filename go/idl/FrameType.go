// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package idl

import "strconv"

type FrameType int16

const (
	FrameTypeMessage    FrameType = 0
	FrameTypeMessageAck FrameType = 1
	FrameTypeOnLine     FrameType = 100
	FrameTypeOnLineAck  FrameType = 101
)

var EnumNamesFrameType = map[FrameType]string{
	FrameTypeMessage:    "Message",
	FrameTypeMessageAck: "MessageAck",
	FrameTypeOnLine:     "OnLine",
	FrameTypeOnLineAck:  "OnLineAck",
}

var EnumValuesFrameType = map[string]FrameType{
	"Message":    FrameTypeMessage,
	"MessageAck": FrameTypeMessageAck,
	"OnLine":     FrameTypeOnLine,
	"OnLineAck":  FrameTypeOnLineAck,
}

func (v FrameType) String() string {
	if s, ok := EnumNamesFrameType[v]; ok {
		return s
	}
	return "FrameType(" + strconv.FormatInt(int64(v), 10) + ")"
}