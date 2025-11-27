package stdmsg

import (
	"time"
)

type Rfq struct {
	Header       MessageHeader
	InstrumentId uint32
	State        YesNoEnum
	Side         RfqDirectionEnum
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
	// m.State.Decode(c)
	// m.Side.Decode(c)
	c.Decode(&m.Amount)
	c.Decode(&m.TimestampMs)
}

type RfqDirectionEnum uint8
type RfqDirectionValues struct {
	Buy          RfqDirectionEnum
	Sell         RfqDirectionEnum
	No_direction RfqDirectionEnum
	NullValue    RfqDirectionEnum
}

var RfqDirection = RfqDirectionValues{0, 1, 2, 255}

func (m *RfqDirectionEnum) PPrint(i int) {
	switch *m {
	case RfqDirection.Buy:
		PPrintlnInd(i, "Side: buy")
	case RfqDirection.Sell:
		PPrintlnInd(i, "Side: sell")
	case RfqDirection.No_direction:
		PPrintlnInd(i, "Side: no direction")
	default:
		PPrintlnInd(i, "Side: null")
	}
}

func (m *RfqDirectionEnum) Decode(c *Coder) {
	c.Decode(m)
}
