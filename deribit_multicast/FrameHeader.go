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
	PPrintlnInd(i, "Frame Header")
	PPrintlnInd(i+2, "PacketLength:", h.PacketLength)
	PPrintlnInd(i+2, "ChainId:", h.ChainId)
	PPrintlnInd(i+2, "SequenceNumber:", h.SequenceNumber)
}

func (h *FrameHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	_m.b[0] = byte(h.PacketLength)
	_m.b[1] = byte(h.PacketLength >> 8)
	_m.b[2] = byte(h.ChainId)
	_m.b[3] = byte(h.ChainId >> 8)
	_m.b[4] = byte(h.SequenceNumber)
	_m.b[5] = byte(h.SequenceNumber >> 8)
	_m.b[6] = byte(h.SequenceNumber >> 16)
	_m.b[7] = byte(h.SequenceNumber >> 24)
	if _, err := _w.Write(_m.b); err != nil {
		return err
	}
	return nil
}

func (h *FrameHeader) Decode(_m *SbeGoMarshaller, _r io.Reader) error {
	if _, err := io.ReadFull(_r, _m.b); err != nil {
		return err
	}
	// fmt.Println(">>>", _m)
	h.PacketLength = uint16(_m.b[0]) | uint16(_m.b[1])<<8
	h.ChainId = uint16(_m.b[2]) | uint16(_m.b[3])<<8
	h.SequenceNumber = (uint32(_m.b[4]) | uint32(_m.b[5])<<8 |
		uint32(_m.b[6])<<16 | uint32(_m.b[7])<<24)
	return nil
}

func (h *FrameHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}
