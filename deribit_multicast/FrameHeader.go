// Generated SBE (Simple Binary Encoding) message codec

package deribit_multicast

import (
	"io"
)

type FrameHeader struct {
	PacketLength   uint16
	ChainId        uint16
	SequenceNumber uint32
}

func (h *FrameHeader) PPrint(i int) {
	PPrintlnInd(i, "Object: Frame Header")
	PPrintlnInd(i+2, "PacketLength:", h.PacketLength)
	PPrintlnInd(i+2, "ChainId:", h.ChainId)
	PPrintlnInd(i+2, "SequenceNumber:", h.SequenceNumber)
}

func (m *FrameHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	_m.b8[0] = byte(m.PacketLength)
	_m.b8[1] = byte(m.PacketLength >> 8)
	_m.b8[2] = byte(m.ChainId)
	_m.b8[3] = byte(m.ChainId >> 8)
	_m.b8[4] = byte(m.SequenceNumber)
	_m.b8[5] = byte(m.SequenceNumber >> 8)
	_m.b8[6] = byte(m.SequenceNumber >> 16)
	_m.b8[7] = byte(m.SequenceNumber >> 24)
	if _, err := _w.Write(_m.b8); err != nil {
		return err
	}
	return nil
}

func (m *FrameHeader) Decode(_m *SbeGoMarshaller, _r io.Reader) error {
	if _, err := io.ReadFull(_r, _m.b8); err != nil {
		return err
	}
	m.PacketLength = uint16(_m.b8[0]) | uint16(_m.b8[1])<<8
	m.ChainId = uint16(_m.b8[2]) | uint16(_m.b8[3])<<8
	m.SequenceNumber = (uint32(_m.b8[4]) | uint32(_m.b8[5])<<8 |
		uint32(_m.b8[6])<<16 | uint32(_m.b8[7])<<24)
	return nil
}

func (m *FrameHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}
