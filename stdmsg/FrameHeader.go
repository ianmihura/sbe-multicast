package stdmsg

type FrameHeader struct {
	PacketLength   uint16
	ChannelId      uint16
	SequenceNumber uint32
}

func (m *FrameHeader) PPrint(i int) {
	PPrintlnInd(i, "Frame Header")
	PPrintlnInd(i+2, "PacketLength:", m.PacketLength)
	PPrintlnInd(i+2, "ChannelId:", m.ChannelId)
	PPrintlnInd(i+2, "SequenceNumber:", m.SequenceNumber)
}

func (m *FrameHeader) Decode(c *Coder) {
	c.Decode(&m.PacketLength)
	c.Decode(&m.ChannelId)
	c.Decode(&m.SequenceNumber)
}
