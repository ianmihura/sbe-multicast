package stdmsg

import (
	"time"
)

type Rfq struct {
	Header       MessageHeader
	InstrumentId uint32
	State        YesNoEnum
	Side         DirectionEnum
	Amount       float64
	TimestampMs  uint64
}

func (m *Rfq) PPrint(i int) {
	PPrintlnInd(i, "RFQ")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", m.InstrumentId)
	m.State.PPrintCustom(i+2, "State:", "inactive", "active")
	m.Side.PPrint(i + 2)
	PPrintlnInd(i+2, "Amount:", m.Amount)
	PPrintlnInd(i+2, "TimestampMs:", time.UnixMilli(int64(m.TimestampMs)))
}

func (m *Rfq) Decode(c *Coder) {
	c.Decode(&m.InstrumentId)
	c.Decode(&m.State)
	c.Decode(&m.Side)
	c.Decode(&m.Amount)
	c.Decode(&m.TimestampMs)
}
