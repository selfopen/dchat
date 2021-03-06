// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package idl

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OnLineAck struct {
	_tab flatbuffers.Table
}

func GetRootAsOnLineAck(buf []byte, offset flatbuffers.UOffsetT) *OnLineAck {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OnLineAck{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsOnLineAck(buf []byte, offset flatbuffers.UOffsetT) *OnLineAck {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OnLineAck{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *OnLineAck) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OnLineAck) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OnLineAck) Header(obj *Header) *Header {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Header)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *OnLineAck) Id(obj *UByte16) *UByte16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(UByte16)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *OnLineAck) Partner(obj *Partner) *Partner {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Partner)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *OnLineAck) Ts(obj *Timestamp) *Timestamp {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Timestamp)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func OnLineAckStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func OnLineAckAddHeader(builder *flatbuffers.Builder, header flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(header), 0)
}
func OnLineAckAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(id), 0)
}
func OnLineAckAddPartner(builder *flatbuffers.Builder, partner flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(partner), 0)
}
func OnLineAckAddTs(builder *flatbuffers.Builder, ts flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(ts), 0)
}
func OnLineAckEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
